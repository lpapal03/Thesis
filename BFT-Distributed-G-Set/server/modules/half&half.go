package modules

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
)

func Server_Task_HalfAndHalf(me config.Node, peers []config.Node) {

	server := server.Create(me, peers)

	// Listen to messages and handle them
	for {
		message, _ := server.Receive_socket.RecvMessage(0)
		messaging.HandleMessage(server, message)
	}
}
