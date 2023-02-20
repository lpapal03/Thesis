package modules

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
	"backend/tools"

	zmq "github.com/pebbe/zmq4"
)

func Mute_listener_task(listener config.Node, peers []config.Node, zctx *zmq.Context) {
	server := server.CreateServer(listener, peers, zctx)
	for {
		msg, _ := server.Receive_socket.RecvMessage(0)
		message, err := messaging.ParseMessageString(msg)
		if err != nil {
			continue
		}
		tools.Log(server.Id, "Received "+message.Tag+" from "+message.Sender+", no action")
	}
}
