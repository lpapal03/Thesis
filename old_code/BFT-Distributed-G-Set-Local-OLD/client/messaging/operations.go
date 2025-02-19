package messaging

import (
	"frontend/client"
	"frontend/config"
	"frontend/tools"
	"math/rand"
	"reflect"
	"sort"
	"strings"

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

// returns true iff we have
// more than 2f+1 replies
// and f+1 matching
// returns a valid reply
func findGetValidReply(replies map[string]string) string {
	if len(replies) < 2*config.F+1 {
		return ""
	}
	reply_strings := []string{}
	for k := range replies {
		temp := strings.Split(replies[k], ",")
		sort.Strings(temp)
		reply_strings = append(reply_strings, strings.Join(temp, " "))
	}

	histo := make(map[string]int)
	for _, str := range reply_strings {
		histo[str]++
	}
	valhi := 0
	strhi := ""
	for k, v := range histo {
		if v > valhi {
			valhi = v
			strhi = k
		}
	}
	if valhi >= config.F+1 {
		return strhi
	}
	return ""
}

func Get(c *client.Client) string {
	tools.Log(c.Id, "Called GET")
	c.Message_counter++
	sendToServers(c.Servers, []string{GET}, 3*config.F+1)

	replies := make(map[string]string)
	tools.Log(c.Id, "Waiting for valid GET_REPLY")
	for {
		sockets, _ := c.Poller.Poll(-1)
		for _, socket := range sockets {
			s := socket.Socket
			msg, _ := s.RecvMessage(0)
			if msg[1] == GET_RESPONSE {
				replies[msg[0]] = msg[2]
			}
		}
		r := findGetValidReply(replies)
		if len(r) > 0 {
			tools.Log(c.Id, "Reply: "+r)
			return r
		}
	}
}

// TODO: Handle responses
// Do i have to send to 2f+1 or all?
func Add(c *client.Client, record string) {
	tools.Log(c.Id, "Called ADD("+record+")")
	sendToServers(c.Servers, []string{ADD, record}, 2*config.F+1)
	// WAIT FOR F+1 RESPONSES
	replies := make(map[string]bool)
	tools.Log(c.Id, "Waiting for f+1 ADD replies")
	for {
		sockets, _ := c.Poller.Poll(-1)
		for _, socket := range sockets {
			s := socket.Socket
			msg, _ := s.RecvMessage(0)
			if msg[1] == ADD_RESPONSE && msg[2] == record {
				replies[msg[0]] = true
			}
		}
		if len(replies) >= config.F+1 {
			tools.Log(c.Id, "Record {"+record+"} appended")
			return
		}
	}
}
