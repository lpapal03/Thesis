// Client

package main

import (
	"client/config"
	"fmt"
	"sort"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func broadcast(message string, server_sockets []*zmq.Socket) {
	for i := 0; i < len(server_sockets); i++ {
		server_sockets[i].SendMessage(message)
	}
}

func get(server_sockets []*zmq.Socket, msg_cnt *int, poller *zmq.Poller) {
	*msg_cnt += 1
	broadcast("get", server_sockets)
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
					if msg[0] == "get_response" {
						reply_messages = append(reply_messages, msg[1])
						replies += 1
					}
				}
			}
		}
	}

	fmt.Println("GET operation done")

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
		fmt.Println(reply_messages[i])
	}

	// We can now begin comparing server replies
	// In order to find f+1 matching replies
	var matching_replies int = 0
	for i := 0; i < len(reply_messages); i++ {
		matching_replies++
		if matching_replies >= config.LOW_THRESHOLD {
			break
		}
	}

	if matching_replies >= config.LOW_THRESHOLD {
		// return reply that is same in f+1 replies
	}

}

func client_task(id string, server_ports []string) {

	// Declare context, poller, router sockets of servers, message counter
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	var server_sockets []*zmq.Socket
	message_counter := 0

	// Connect client dealer sockets to all servers
	for i := 0; i < len(server_ports); i++ {
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(id)
		s.Connect("tcp://localhost:" + server_ports[i])
		fmt.Println("Client conected to", "tcp://localhost:"+server_ports[i])
		server_sockets = append(server_sockets, s)
		poller.Add(server_sockets[i], zmq.POLLIN)
	}

	get(server_sockets, &message_counter, poller)

}

func main() {

	go client_task("c1", config.Servers)

	for {
	}

}
