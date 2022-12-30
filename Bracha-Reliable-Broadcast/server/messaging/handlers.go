package messaging

import (
	"backend/server"
	"backend/tools"
)

func HandleMessage(server server.Server, msg []string) {
	message, err := StringToMessage(msg)
	// if message is not in the right format, discard
	if err != nil {
		return
	}
	tools.Log(server.Id, "Received "+message.Tag+" from "+message.Sender)

	if message.Tag == BRACHA_BROADCAST {
		handleRB(server, message)
	} else {
		handleRB_response(server, message)
	}

}

func handleRB(receiver server.Server, message Message) {
	ReliableBroadcast(receiver, message)
}

func handleRB_response(receiver server.Server, message Message) {
	delivered := HandleReliableBroadcast(receiver, message)
	if delivered {
		tools.Log(receiver.Id, "Delivered {"+message.Content[1]+"}")
	}
}
