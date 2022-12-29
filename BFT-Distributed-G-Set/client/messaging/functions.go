package messaging

import (
	"frontend/client"
	"frontend/config"
	"frontend/tools"
	"sort"
	"strings"
)

// returns true iff we have
// more than 2f+1 replies
// and f+1 matching
// returns a valid reply
func findValidReply(replies map[string]string) string {
	if len(replies) < config.MEDIUM_THRESHOLD {
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
	if valhi >= config.LOW_THRESHOLD {
		return strhi
	}
	return ""
}

func Get(client client.Client) string {

	for i := 0; i < len(client.Servers); i++ {
		client.Servers[i].SendMessage([]string{GET})
	}

	tools.Log(client.Id, "Called GET")
	client.Message_counter++
	replies := make(map[string]string)

	tools.Log(client.Id, "Waiting for valid GET_REPLY...")
	for {
		sockets, _ := client.Poller.Poll(-1)
		for _, socket := range sockets {
			s := socket.Socket
			msg, _ := s.RecvMessage(0)
			if msg[1] == GET_RESPONSE {
				replies[msg[0]] = msg[2]
			}
		}
		r := findValidReply(replies)
		if len(r) > 0 {
			tools.Log(client.Id, "Reply: "+r)
			return r
		}
	}
}

// TODO: Handle responses
func Add(client client.Client, record string) {
	tools.Log(client.Id, "Called ADD("+record+")")
	for i := 0; i < len(client.Servers); i++ {
		// TargetedAdd(client, *client.Servers[i], record)
		client.Message_counter++
		client.Servers[i].SendMessage([]string{ADD, record})
	}
}
