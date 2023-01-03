package scenarios

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func Normal_listener_task(listener config.Node, peers []config.Node) {
	server := server.Create(listener, peers)
	for {
		message, err := server.Receive_socket.RecvMessage(0)
		if err != nil {
			server.Receive_socket.Close()
			new_sock, _ := server.Zctx.NewSocket(zmq.ROUTER)
			server.Receive_socket = *new_sock
			fmt.Println("Renewed socket")
		}
		messaging.HandleMessage(server, message)
	}
}
