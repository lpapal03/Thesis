package messaging

import (
	"backend/gset"
	"backend/server"
	"backend/tools"
	"strings"
)

func HandleMessage(server server.Server, message Message) {
	tools.Log(server.Id, "Received "+message.Tag+" from "+message.Sender)

	if message.Tag == GET {
		handleGet(server, message)
	}
	if message.Tag == ADD {
		handleAdd(server, message)
	}
	// message is related to broadcast service
	if strings.Contains(message.Tag, BRACHA_BROADCAST) {
		HandleReliableBroadcast(server, message)
	}
}

// Handle get request. I need sender_id to know where
// my response will go to
func handleGet(server server.Server, message Message) {
	response := []string{message.Sender, server.Id, GET_RESPONSE, gset.GsetToString(server.Gset, false)}
	server.Receive_socket.SendMessage(response)
	tools.Log(server.Id, GET_RESPONSE+" to "+message.Sender)
}

func handleAdd(server server.Server, message Message) {

	// Call RB service
	ReliableBroadcast(server, message)
	// if true, append

}
