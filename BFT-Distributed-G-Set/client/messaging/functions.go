package messaging

import (
	"frontend/client"
	"frontend/config"
	"frontend/tools"
	"math/rand"
	"sort"
	"strings"
	"time"
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

// the amount of add responses given a record
func countAddReplies(replies map[string]bool, record string) int {
	count := 0
	for k := range replies {
		if record == strings.Split(k, "-")[1] {
			count++
		}
	}
	return count
}

// TODO: Handle responses
func Add(client client.Client, record string) {
	tools.Log(client.Id, "Called ADD("+record+")")
	//randomly choose 2f+1 servers to send add
	rand.Seed(time.Now().Unix())
	receiver_indexes := rand.Perm(len(client.Servers))
	for i := 0; i < 2*config.F+1; i++ {
		client.Servers[receiver_indexes[i]].SendMessage([]string{ADD, record})
	}
	// client.Servers[0].SendMessage([]string{ADD, record})
	// WAIT FOR F+1 RESPONSES
	// replies := make(map[string]bool)
	// tools.Log(client.Id, "Waiting for f+1 ADD_RESPONSES...")
	// for {
	// 	sockets, _ := client.Poller.Poll(-1)
	// 	for _, socket := range sockets {
	// 		s := socket.Socket
	// 		msg, _ := s.RecvMessage(0)
	// 		fmt.Println(msg)
	// 		if msg[1] == ADD_RESPONSE {
	// 			replies[msg[0]+"-"+msg[2]] = true
	// 		}
	// 	}
	// 	if countAddReplies(replies, record) >= config.F+1 {
	// 		tools.Log(client.Id, "Record appended")
	// 		return
	// 	}
	// }
}
