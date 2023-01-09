package scenarios

import (
	"backend/config"
	"backend/messaging"
	"backend/server"

	zmq "github.com/pebbe/zmq4"
)

func Normal_listener_task(listener config.Node, peers []config.Node, zctx *zmq.Context) {
	server := server.CreateServer(listener, peers, zctx)
	for {
		message, err := server.Receive_socket.RecvMessage(0)
		if err != nil {
			panic(err)
		}
		messaging.HandleMessage(server, message)
	}
}
