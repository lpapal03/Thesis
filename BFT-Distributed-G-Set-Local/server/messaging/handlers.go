package messaging

import (
	"backend/gset"
	"backend/server"
	"backend/tools"
	"strings"
)

func HandleMessage(s *server.Server, msg []string) {
	message, err := ParseMessageString(msg)
	if err != nil {
		panic(err)
	}
	tools.Log(s.Id, "Received "+message.Tag+" "+strings.Join(message.Content, ".")+" from "+message.Sender)

	if message.Tag == GET {
		handleGet(s, message)
	} else if message.Tag == ADD {
		handleAdd(s, message)
	} else if strings.Contains(message.Tag, BRACHA_BROADCAST) {
		handleRB(s, message)
	}

}

// Handle get request. I need sender_id to know where
// my response will go to
func handleGet(receiver *server.Server, message Message) {
	response := []string{message.Sender, receiver.Id, GET_RESPONSE, gset.GsetToString(receiver.Gset, false)}
	receiver.Receive_socket.SendMessage(response)
	tools.Log(receiver.Id, GET_RESPONSE+" to "+message.Sender)
}

func handleAdd(receiver *server.Server, message Message) {
	if !gset.Exists(receiver.Gset, message.Content[0]) {
		ReliableBroadcast(receiver, message)
	} else {
		response := []string{message.Sender, receiver.Id, ADD_RESPONSE, message.Content[0]}
		receiver.Receive_socket.SendMessage(response)
	}
}

func handleRB(receiver *server.Server, message Message) {
	response := []string{message.Content[0], receiver.Id, ADD_RESPONSE, message.Content[1]}

	if gset.Exists(receiver.Gset, message.Content[1]) {
		receiver.Receive_socket.SendMessage(response)
		return
	}

	delivered := HandleReliableBroadcast(receiver, message)

	if delivered && !gset.Exists(receiver.Gset, message.Content[1]) {
		gset.Append(receiver.Gset, message.Content[1])
		receiver.Receive_socket.SendMessage(response)
		tools.Log(receiver.Id, "Appended record "+message.Content[1])
		return
	}
	if delivered && gset.Exists(receiver.Gset, message.Content[1]) {
		receiver.Receive_socket.SendMessage(response)
		tools.Log(receiver.Id, "Record "+message.Content[1]+" already exists")
		return
	}
}
