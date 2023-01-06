package scenarios

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
	"backend/tools"
)

func Mute_listener_task(listener config.Node, peers []config.Node) {
	server := server.Create(listener, peers)
	for {
		msg, _ := server.Receive_socket.RecvMessage(0)
		message, err := messaging.ParseMessageString(msg)
		if err != nil {
			continue
		}
		tools.Log(server.Id, "Received "+message.Tag+" from "+message.Sender+", no action")
	}
}
