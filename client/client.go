// Client

package main

import (
	"client/config"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func broadcast(message string, server_sockets []*zmq.Socket) {
	for i := 0; i < len(server_sockets); i++ {
		server_sockets[i].SendMessage(message)
	}
}

func get(server_sockets []*zmq.Socket, msg_cnt int) {
	msg_cnt += 1
	broadcast("get", server_sockets)
	// wait for 2f+1 replies
	// Wait for replies code
	// outside for will happen until i get 2f+1 replies
	// for {
	// 	poller_sockets, _ := poller.Poll(-1)
	// 	for _, poller_socket := range poller_sockets {
	// 		p_s := poller_socket.Socket
	// 		for _, server_socket := range server_sockets {
	// 			if server_socket == p_s {
	// 				msg, _ := p_s.Recv(0)
	// 				fmt.Println(msg)
	// 			}
	// 		}
	// 	}
	// }
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

	get(server_sockets, message_counter)

}

func main() {

	go client_task("c1", config.Servers)

	for {
	}

}
