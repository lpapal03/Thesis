package messaging

import (
	"backend/gset"
	"backend/server"
	"backend/tools"
)

func HandleMessage(server server.Server, message []string) {
	sender_id := message[0]
	message_type := message[1]
	tools.Log(server.Id, "Received "+message_type+" from "+sender_id)

	if message_type == GET {
		handleGet(server, sender_id)
	}
	if message_type == ADD {
		handleAdd(server, message)
	}
	// message is related to broadcast service
	// if strings.Contains(message_type, BRACHA_BROADCAST) {
	// 	HandleReliableBroadcast(my_id, msg, poller, receive_socket, &echo, &vote)
	// }
}

func handleGet(server server.Server, sender_id string) {
	response := []string{sender_id, server.Id, GET_RESPONSE, gset.GsetToString(server.Gset, false)}
	server.Receive_socket.SendMessage(response)
	tools.Log(server.Id, GET_RESPONSE+" to "+sender_id)
}

func handleAdd(server server.Server, message []string) {

	// Call RB service
	ReliableBroadcast()
	// if true, append

}
