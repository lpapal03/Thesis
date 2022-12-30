package modules

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
)

func Normal_listener_task(listener config.Node, peers []config.Node) {
	server := server.Create(listener, peers)
	for {
		message, _ := server.Receive_socket.RecvMessage(0)
		messaging.HandleMessage(server, message)
	}
}
