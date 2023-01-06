package messaging

import (
	"BFT-Distributed-G-Set/gset"
	"BFT-Distributed-G-Set/server"
	"BFT-Distributed-G-Set/tools"
	"fmt"
	"strings"
)

func HandleMessage(server server.Server, msg []string) {
	message, err := ParseMessageString(msg)
	if err != nil {
		fmt.Println(err)
	}
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
func handleGet(receiver server.Server, message Message) {
	response := []string{message.Sender, receiver.Id, GET_RESPONSE, gset.GsetToString(receiver.Gset, false)}
	receiver.Receive_socket.SendMessage(response)
	tools.Log(receiver.Id, GET_RESPONSE+" to "+message.Sender)
}

func handleAdd(receiver server.Server, message Message) {
	if !gset.Exists(receiver.Gset, message.Content[0]) {
		ReliableBroadcast(receiver, message)
	}
}

func handleRB(receiver server.Server, message Message) {
	// original_sender := message.Content[0]
	// response := []string{original_sender, receiver.Id, ADD_RESPONSE, message.Content[1]}

	delivered := HandleReliableBroadcast(receiver, message)
	if delivered {
		gset.Append(receiver.Gset, message.Content[1])
		// receiver.Receive_socket.SendMessage(response)
		tools.Log(receiver.Id, "Appended record!")
	}
}
