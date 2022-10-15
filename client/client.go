// Client

package main

import (
	"fmt"
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"
)

// Sender id is bound to the socket
// func get(s *zmq.Socket, msg_id int) {
// 	msg := []string{strconv.Itoa(msg_id), "get"}
// 	s.SendMessage(msg)
// 	rec_msg, _ := s.RecvMessage(0)
// 	fmt.Println("Server response:\n-------")
// 	fmt.Println(rec_msg[1])

// Send request to 3f+1 servers

// Wait response from 2f+1
// When waiting for responses,

// Foreach record r in each response, r should be in at least f+1 responses
// In other words, f+1 responses should match
// }
func client_task(id string) {

	rand.Seed(time.Now().UnixNano())
	servers := []string{"tcp://localhost:5555", "tcp://localhost:5556"}
	var server_sockets [2]*zmq.Socket
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()

	// id := strconv.Itoa(os.Getpid()) + strconv.Itoa(rand.Intn(10))

	for i := 0; i < len(servers); i++ {
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(id)
		s.Connect(servers[i])
		server_sockets[i] = s
		poller.Add(server_sockets[i], zmq.POLLIN)
		fmt.Printf("Client with id %s connected to server %s\n", id, servers[i])
		fmt.Println(server_sockets[i])
		fmt.Println()
	}

	server_sockets[1].SendMessage("Hello 1")
	server_sockets[0].SendMessage("Hello 0")

	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch s := socket.Socket; s {
			case server_sockets[0]:
				msg, _ := s.Recv(0)
				//  Process msg
				fmt.Println(msg)
			case server_sockets[1]:
				msg, _ := s.Recv(0)
				//  Process msg
				fmt.Println(msg)
			}
		}
	}

}

func main() {

	go client_task("c1")
	// go client_task("c2")
	// go client_task("c3")

	for {
	}

}
