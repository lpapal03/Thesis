package gset

import (
	"client/config"
	"client/messaging"
	"client/tools"
	"errors"
	"sort"
	"strconv"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func Get(me string, server_sockets []*zmq.Socket, msg_cnt *int, poller *zmq.Poller) (string, error) {
	tools.Log(me, "Invoked GET")
	*msg_cnt += 1
	message := []string{messaging.GET}
	messaging.SimpleBroadcast(message, server_sockets)
	// Wait for 2f+1 replies
	var reply_messages = []string{}
	var replies int = 0
	for replies < config.MEDIUM_THRESHOLD {
		poller_sockets, _ := poller.Poll(-1)
		for _, poller_socket := range poller_sockets {
			p_s := poller_socket.Socket
			for _, server_socket := range server_sockets {
				if server_socket == p_s {
					msg, _ := p_s.RecvMessage(0)
					// msg[1] = msg_type
					if msg[1] == messaging.GET_RESPONSE {
						tools.Log(me, "GET response from "+msg[0])
						reply_messages = append(reply_messages, msg[2])
						replies += 1
					}
				}
			}
		}
	}

	tools.Log(me, messaging.GET+" done, received "+strconv.Itoa(len(reply_messages))+"/"+strconv.Itoa(config.LOW_THRESHOLD)+" wanted replies")

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
				tools.Log(me, "Found "+strconv.Itoa(matching_replies)+"/"+strconv.Itoa(config.LOW_THRESHOLD)+" matching replies")
				return reply_messages[i], nil
			}
		}
	}
	return "", errors.New("No f+1 matching responses!")
}

func Add(me string, server_sockets []*zmq.Socket, msg_cnt *int, poller *zmq.Poller, record string) {
	tools.Log(me, "Invoked ADD with {"+record+"}")
	*msg_cnt += 1
	messaging.SimpleBroadcast([]string{messaging.ADD, record}, server_sockets)
}
