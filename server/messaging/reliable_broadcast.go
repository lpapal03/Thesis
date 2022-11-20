package messaging

import (
	"backend/server"
	"backend/tools"
	"fmt"
	"strings"
)

// Implement reliable broadcast
// Here, all of the sending and receiving will happen
// CALLED ONCE BY LEADER
func ReliableBroadcast(leader server.Server, message Message) {

	// step 0
	server.CreateBroadcastState(leader, message.Content)
	msg := CreateMessage(BRACHA_BROADCAST_INIT, append([]string{message.Sender}, message.Content...))
	for _, pier_sokcet := range leader.Piers {
		pier_sokcet.SendMessage(msg)
	}
	tools.Log(leader.Id, "Sent message {"+(strings.Join(msg, " "))+"} to all piers")
	// message format: [localhost:10000 BRACHA_BROADCAST_INIT c1 ADD hello]

}

// Called from every server receiving RB messages
func HandleReliableBroadcast(receiver server.Server, message Message) {

	// Message comes. I add it to the correct spot
	// Then I check the current step of the server
	// Then I act accordingly

	fmt.Println(message.Content)

	message_identifier := strings.Join(message.Content, " ")
	pier_sender := message.Sender

	if message.Tag == BRACHA_BROADCAST_INIT {
		// This is an entry in the map of echo and read values
		server.CreateBroadcastState(receiver, message.Content)
		sendToAll(receiver, message, BRACHA_BROADCAST_ECHO)

	} else if message.Tag == BRACHA_BROADCAST_ECHO {
		fmt.Println(message_identifier)
		// receiver.BroadcastState[message_identifier].Echo[pier_sender] = true

	} else if message.Tag == BRACHA_BROADCAST_READY {
		receiver.BroadcastState[message_identifier].Ready[pier_sender] = true

	}

}

func sendToAll(receiver server.Server, message Message, new_tag string) {
	msg := CreateMessage(new_tag, append([]string{message.Sender}, message.Content...))
	for _, pier_sokcet := range receiver.Piers {
		pier_sokcet.SendMessage(msg)
	}
}
