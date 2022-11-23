package messaging

import (
	"backend/gset"
	"backend/server"
	"backend/tools"
	"strings"
)

func HandleMessage(server server.Server, msg []string) {
	message := ParseMessage(msg)
	tools.Log(server.Id, "Received "+message.Tag+" from "+message.Sender)

	if message.Tag == GET {
		handleGet(server, message)
	} else if message.Tag == ADD {
		handleAdd(server, message)
	} else if strings.Contains(message.Tag, BRACHA_BROADCAST) {
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
	if !gset.Exists(server.Gset, message.Content[0]) {
		ReliableBroadcast(server, message)
	} else {
		response := []string{message.Sender, server.Id, ADD_RESPONSE, "Success", message.Content[0]}
		server.Receive_socket.SendMessage(response)
		tools.Log(server.Id, "Sent ADD_RESPONSE to "+message.Sender)
	}
}

// called when RB is done
func handleAddInternal(server server.Server, message Message) {

	gset.Append(server.Gset, message.Content[1])
	response := []string{message.Content[0], server.Id, ADD_RESPONSE, "Success", message.Content[1]}
	server.Receive_socket.SendMessage(response)

	tools.Log(server.Id, "Sent ADD_RESPONSE to "+message.Sender)
}

func handleRB(receiver server.Server, message Message) {
	HandleReliableBroadcast(receiver, message)
}
