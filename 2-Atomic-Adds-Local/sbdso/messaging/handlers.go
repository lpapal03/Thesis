package messaging

import (
	"backend/config"
	"backend/gset"
	"backend/server"
	"backend/tools"
	"math/rand"
	"reflect"
	"strings"

	"github.com/pebbe/zmq4"
)

func HandleMessage(s *server.Server, msg []string) {
	message, err := ParseMessageString(msg)
	if err != nil {
		panic(err)
	}
	if message.Tag == GET {
		tools.Log(s.Id, "Received "+message.Tag+"from "+message.Sender)
	} else {
		tools.Log(s.Id, "Received "+message.Tag+" {"+strings.Join(message.Content, " ")+"} from "+message.Sender)
	}
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

		gset.Add(receiver.Gset, message.Content[1])
		msg_signatue := strings.Split(message.Content[1], ";")[2]
		r1, r2 := gset.CheckAtomic(receiver.Gset, msg_signatue)
		if len(r1) > 0 && len(r2) > 0 {
			handleAtomicAdd(receiver, r1, r2)
		}

		receiver.Receive_socket.SendMessage(response)
		tools.Log(receiver.Id, "Appended record {"+message.Content[1]+"}")
		return
	}
	if delivered && gset.Exists(receiver.Gset, message.Content[1]) {
		receiver.Receive_socket.SendMessage(response)
		tools.Log(receiver.Id, "Record {"+message.Content[1]+"} already exists")
		return
	}
}

func handleAtomicAdd(s *server.Server, r1, r2 string) {
	tools.Log(s.Id, "Found atomic records {"+r1+"} with {"+r2+"}")
	var response []string

	// handle
	parts1, parts2 := strings.Split(r1, ";"), strings.Split(r2, ";")
	client1, client2 := parts1[1], parts2[1]
	dest1, dest2 := parts1[3], parts2[3]
	msg1, msg2 := parts1[4], parts2[4]

	// handle 1
	response = []string{client1, s.Id, ADD_ATOMIC_RESPONSE, r1}
	s.Receive_socket.SendMessage(response)
	tools.Log(s.Id, "Sent ADD_ATOMIC_RESPONSE to "+client1)
	Add(s, msg1, dest1)

	// handle 2
	response = []string{client2, s.Id, ADD_ATOMIC_RESPONSE, r2}
	s.Receive_socket.SendMessage(response)
	tools.Log(s.Id, "Sent ADD_ATOMIC_RESPONSE to "+client2)
	Add(s, msg2, dest2)

}

func Add(s *server.Server, record, destination string) {
	tools.Log(s.Id, "Called ADD("+record+") with destination:"+destination)
	sendToServers(s.Bdso_networks[destination], []string{ADD, record}, 2*config.F+1)
}

func sendToServers(m map[string]*zmq4.Socket, message []string, amount int) {
	keys := reflect.ValueOf(m).MapKeys()
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
	for i := 0; i < amount; i++ {
		key := keys[i].String()
		s := m[key]
		s.SendMessage(message)
	}
}
