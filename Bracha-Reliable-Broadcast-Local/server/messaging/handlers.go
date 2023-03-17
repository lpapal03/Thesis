package messaging

import (
	"backend/server"
	"backend/tools"
	"strings"
)

func HandleMessage(s *server.Server, msg []string) {
	message, err := ParseMessageString(msg)
	if err != nil {
		return
	}
	tools.Log(s.Id, "Received "+message.Tag+" "+strings.Join(message.Content, ".")+" from "+message.Sender)

	if message.Tag == BRACHA_BROADCAST {
		handleRBInit(s, message)
	} else if strings.Contains(message.Tag, BRACHA_BROADCAST) {
		handleRB(s, message)
	}
}

func handleRBInit(receiver *server.Server, message Message) {
	ReliableBroadcast(receiver, message)
}

func handleRB(receiver *server.Server, message Message) {

	HandleReliableBroadcast(receiver, message)

}
