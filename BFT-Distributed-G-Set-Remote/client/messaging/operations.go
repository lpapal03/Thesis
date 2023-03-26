package messaging

import (
	"2-Atomic-Adds/client"
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/tools"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pebbe/zmq4"
)

func sendToServers(m map[string]*zmq4.Socket, message []string, amount int) {
	keys := reflect.ValueOf(m).MapKeys()
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
	for i := 0; i < amount; i++ {
		key := keys[i].String()
		s := m[key]
		s.SendMessage(message)
	}
}

func countMatchingReplies(replies map[string]string) string {
	if len(replies) < 2*config.F+1 {
		return ""
	}
	// Create a map to count the occurrences of each record
	counts := make(map[string]int)
	for _, reply := range replies {
		// Parse the reply into individual records
		records := strings.Split(reply, ",")
		for _, record := range records {
			// Increment the count for this record
			counts[record]++
		}
	}

	// Find the records that appear in at least F+1 of the sets
	var commonSet []string
	for record, count := range counts {
		if count >= config.F+1 {
			commonSet = append(commonSet, record)
		}
	}

	return strings.Join(commonSet, " ")
}

func Get(c *client.Client) (string, time.Duration) {
	tools.Log(c.Id, "Called GET")
	c.Message_counter++
	start := time.Now()
	sendToServers(c.Servers, []string{GET, strconv.Itoa(c.Message_counter)}, 3*config.F+1)

	replies := make(map[string]string)
	tools.Log(c.Id, "Waiting for valid GET_REPLY")
	for {
		sockets, _ := c.Poller.Poll(-1)
		for _, socket := range sockets {
			s := socket.Socket
			msg, _ := s.RecvMessage(0)
			tools.Log(c.Id, strings.Join(msg, "-"))
			if msg[1] == GET_RESPONSE {
				replies[msg[0]] = msg[2]
			}
		}
		r := countMatchingReplies(replies)
		if len(r) > 0 {
			elapsed := time.Since(start)
			tools.Log(c.Id, "GET completed in: "+elapsed.String())
			return r, elapsed
		}
	}
}

// Do i have to send to 2f+1 or all?
func Add(c *client.Client, record string) time.Duration {
	c.Message_counter++
	tools.Log(c.Id, "Called ADD("+record+")")
	message := strconv.Itoa(c.Message_counter) + "." + record
	start := time.Now()
	sendToServers(c.Servers, []string{ADD, message}, 2*config.F+1)
	// WAIT FOR F+1 RESPONSES
	replies := make(map[string]bool)
	tools.Log(c.Id, "Waiting for f+1 ADD replies")
	for {
		sockets, _ := c.Poller.Poll(-1)
		for _, socket := range sockets {
			s := socket.Socket
			msg, _ := s.RecvMessage(0)
			// fmt.Println(msg)
			if strings.Contains(msg[2], ".") {
				msg[2] = strings.Split(msg[2], ".")[2]
			}
			if msg[1] == ADD_RESPONSE && msg[2] == record {
				replies[msg[0]] = true
			}
		}
		if len(replies) >= config.F+1 {
			elapsed := time.Since(start)
			tools.Log(c.Id, "ADD completed in: "+elapsed.String()+". Record {"+record+"} appended")
			return elapsed
		}
	}
}
