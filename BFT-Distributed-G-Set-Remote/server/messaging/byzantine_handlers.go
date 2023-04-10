package messaging

import (
	"BFT-Distributed-G-Set-Remote/server"
	"BFT-Distributed-G-Set-Remote/tools"
	"regexp"
	"strconv"
	"strings"
)

func HandleMessageByzantine(s *server.Server, msg []string, scenario string) {
	message, err := ParseMessageString(msg)
	if err != nil {
		tools.Log(s.Id, "Error msg: "+strings.Join(msg, " "))
		return
	}
	if message.Tag == GET {
		tools.Log(s.Id, "Received "+message.Tag+" from "+message.Sender)
	} else {
		tools.Log(s.Id, "Received "+message.Tag+" {"+strings.Join(message.Content, " ")+"} from "+message.Sender)
	}

	var half_and_half bool

	if scenario == "MALICIOUS" {
		half_and_half = false
	}
	if scenario == "HALF_AND_HALF" {
		half_and_half = true
	}

	// handle
	if message.Tag == GET {
		tools.IncrementNormalCount()
		handleGetByzantine(s, message, half_and_half)
	} else if message.Tag == ADD {
		tools.IncrementNormalCount()
		message.Content[0] = message.Sender + "." + message.Content[0]
		handleAddByzantine(s, message, half_and_half)
	} else if strings.Contains(message.Tag, BRACHA_BROADCAST) {
		tools.IncrementBRBCount()
		handleRBByzantine(s, message, half_and_half)
	}

}

// Handle get request. I need sender_id to know where
// my response will go to
func handleGetByzantine(receiver *server.Server, message Message, half_and_half bool) {
	byzantine_value := generateByzantineValue(message.Sender, half_and_half)
	response := []string{message.Sender, receiver.Id, GET_RESPONSE, byzantine_value}
	receiver.Receive_socket.SendMessage(response)
}

func handleAddByzantine(receiver *server.Server, message Message, half_and_half bool) {
	byzantine_value := generateByzantineValue(message.Sender, half_and_half)
	response := []string{message.Sender, receiver.Id, ADD_RESPONSE, byzantine_value}
	receiver.Receive_socket.SendMessage(response)
}

func handleRBByzantine(receiver *server.Server, message Message, half_and_half bool) {
	byzantine_value := generateByzantineValue(message.Sender, half_and_half)
	var tag string
	switch message.Tag {
	case BRACHA_BROADCAST_INIT:
		tag = BRACHA_BROADCAST_ECHO
	case BRACHA_BROADCAST_ECHO:
		tag = BRACHA_BROADCAST_VOTE
	default:
		tag = ADD_RESPONSE
	}

	msg_parts := strings.Split(message.Content[1], ".")
	msg_parts[2] = byzantine_value
	message.Content[1] = strings.Join(msg_parts, ".")
	v := CreateMessageString(tag, message.Content)
	sendToAll(receiver, v)
}

func generateByzantineValue(sender string, half_and_half bool) string {
	if !half_and_half {
		return "BYZANTINE_0"
	}

	pattern := `node(\d+)`
	re := regexp.MustCompile(pattern)
	match1 := re.FindStringSubmatch(sender)
	var sender_node int
	if len(match1) >= 2 {
		sender_node, _ = strconv.Atoi(match1[1])
	}

	if sender_node%2 == 0 {
		return "BYZANTINE_0"
	} else {
		return "BYZANTINE_1"
	}
}
