// Client

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	zmq "github.com/pebbe/zmq4"
)

// Sender id is bound to the socket
func get(s *zmq.Socket, msg_id int) {
	msg := []string{strconv.Itoa(msg_id), "get"}
	s.SendMessage(msg)
	rec_msg, _ := s.RecvMessage(0)
	fmt.Println("Server response:\n-------")
	fmt.Println(rec_msg[1])

	// Send request to 3f+1 servers

	// Wait response from 2f+1
	// When waiting for responses,

	// Foreach record r in each response, r should be in at least f+1 responses
	// In other words, f+1 responses should match
}

func client_task(id string) {
	rand.Seed(time.Now().UnixNano())
	zctx, _ := zmq.NewContext()

	// id := strconv.Itoa(os.Getpid()) + strconv.Itoa(rand.Intn(10))

	s, _ := zctx.NewSocket(zmq.DEALER)
	s.SetIdentity(id)
	s.Connect("tcp://localhost:5555")

	fmt.Printf("Client with id %s connected and bound\n", id)

	msg_id := 1
	get(s, msg_id)
	msg_id++

}

func main() {

	go client_task("c1")
	go client_task("c2")
	go client_task("c3")

	for {
	}

}
