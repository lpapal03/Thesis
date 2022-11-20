package messaging

import (
	"backend/gset"
	"backend/server"
	"backend/tools"
)

func HandleMessage(server server.Server, message Message) {
	tools.Log(server.Id, "Received "+message.Tag+" from "+message.Sender)

	switch message.Tag {
	case GET:
		handleGet(server, message)
	case ADD:
		handleAdd(server, message)
	case BRACHA_BROADCAST_INIT:
		handleRB(server, message)
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

func handleRB(receiver server.Server, message Message) {
	HandleReliableBroadcast(receiver, message)
}
