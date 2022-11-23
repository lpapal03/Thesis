package messaging

import (
	"errors"
	"frontend/client"
	"frontend/config"
	"frontend/tools"
	"sort"
	"strconv"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func isRecordValid(r string) error {
	if strings.Contains(r, " ") {
		return errors.New("Message can't have spaces!")
	}
	return nil
}

func simpleBroadcast(message []string, servers []*zmq.Socket) {
	for i := 0; i < len(servers); i++ {
		servers[i].SendMessage(message)
	}
}

func countReplies(m map[string]bool) int {
	count := 0
	for _, msg := range m {
		if msg {
			count++
		}
	}
	return count
}

func GetGset(client client.Client) (string, error) {

	simpleBroadcast([]string{GET}, client.Servers)
	tools.Log(client.Id, "Broadcasted {GET} to all servers")

	client.Message_counter++

	// reply matrix ensures that i dont get
	// the same reply from a server more than once.
	var reply_messages = []string{}
	reply_matrix := make(map[string]bool)

	// Wait for 2f+1 replies
	for len(reply_matrix) < config.MEDIUM_THRESHOLD {
		sockets, _ := client.Poller.Poll(-1)
		for _, socket := range sockets {
			s := socket.Socket
			msg, _ := s.RecvMessage(0)
			if msg[1] == GET_RESPONSE && !reply_matrix[msg[0]] {
				reply_matrix[msg[0]] = true
				tools.Log(client.Id, "GET response from "+msg[0])
				reply_messages = append(reply_messages, msg[2])
			}
		}
	}
	tools.Log(client.Id, GET+" done, received "+strconv.Itoa(len(reply_messages))+"/"+strconv.Itoa(config.MEDIUM_THRESHOLD)+" wanted replies")

	// By this point I have 2f+1 replies
	// Now to check if f+1 are the same

	// We need to make sure the replies are comparable
	// For this, we need to separate records, order them and the join them
	// Therefore creating a single string for each reply, which is easily compared
	for i := 0; i < len(reply_messages); i++ {
		// divide reply to individual records
		records := strings.Split(reply_messages[i], "\n")
		// sort records
		sort.Strings(records)
		reply_messages[i] = strings.Join(records, "")
	}

	// We can now begin comparing server replies
	// In order to find f+1 matching replies
	var matching_replies int = 0
	for i := 0; i < len(reply_messages); i++ {
		matching_replies = 0
		for j := 0; j < len(reply_messages); j++ {
			if i == j {
				continue
			}
			if strings.Contains(reply_messages[i], reply_messages[j]) ||
				strings.Contains(reply_messages[j], reply_messages[i]) {
				matching_replies++
			}
			if matching_replies >= config.LOW_THRESHOLD {
				tools.Log(client.Id, "Found "+strconv.Itoa(matching_replies)+"/"+strconv.Itoa(config.LOW_THRESHOLD)+" matching replies")
				return reply_messages[i], nil
			}
		}
	}
	return "", errors.New("No f+1 matching responses!")
}

// TODO: Handle responses
func Add(client client.Client, record string) {
	err := isRecordValid(record)
	if err != nil {
		tools.LogFatal(client.Id, err.Error())
	}
	tools.Log(client.Id, "Invoked ADD with {"+record+"}")
	client.Message_counter++
	simpleBroadcast([]string{ADD, record}, client.Servers)
}

func TargetedAdd(client client.Client, target zmq.Socket, record string) {
	err := isRecordValid(record)
	if err != nil {
		tools.LogFatal(client.Id, err.Error())
	}
	tools.Log(client.Id, "Invoked targeted ADD with {"+record+"}")
	client.Message_counter++
	target.SendMessage([]string{ADD, record})
}
