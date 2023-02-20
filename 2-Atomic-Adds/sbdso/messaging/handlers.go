package messaging

import (
	"2-Atomic-Adds/gset"
	"2-Atomic-Adds/server"
	"2-Atomic-Adds/tools"
	"math/rand"
	"strconv"
	"strings"

	"github.com/pebbe/zmq4"
)

func HandleMessage(s *server.Server, msg []string) {
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

	// handle
	if message.Tag == GET {
		handleGet(s, message)
	} else if message.Tag == ADD {
		message.Content[0] = message.Sender + "." + message.Content[0]
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
	var tag string
	if !gset.Exists(receiver.Gset, message.Content[0]) {
		ReliableBroadcast(receiver, message)
	} else {
		if strings.Contains(message.Content[0], "atomic;") {
			tag = ADD_ATOMIC_RESPONSE
		} else {
			tag = ADD_RESPONSE
		}
		response := []string{message.Sender, receiver.Id, tag, message.Content[0]}
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
		// now check if atomic
		gset.Add(receiver.Gset, message.Content[1])
		if strings.Contains(message.Content[1], "atomic;") {
			r1, r2 := gset.CheckAtomic(receiver.Gset)
			if len(r1) > 0 && len(r2) > 0 {
				handleAtomicAdd(receiver, r1, r2)
			}
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

	// send adds
	BdsoAdd(s, msg1, msg2, dest1, dest2)

	// respond 1
	response = []string{client1, s.Id, ADD_ATOMIC_RESPONSE, r1}
	s.Receive_socket.SendMessage(response)
	tools.Log(s.Id, "Sent ADD_ATOMIC_RESPONSE to "+client1)

	// respond 2
	response = []string{client2, s.Id, ADD_ATOMIC_RESPONSE, r2}
	s.Receive_socket.SendMessage(response)
	tools.Log(s.Id, "Sent ADD_ATOMIC_RESPONSE to "+client2)

}

// only returns when we know the records were appended
func BdsoAdd(s *server.Server, r1, r2, dest1, dest2 string) {
	tools.Log(s.Id, "Called ADD("+r1+") with destination:"+dest1)
	tools.Log(s.Id, "Called ADD("+r2+") with destination:"+dest2)
	network1, ok1 := s.Bdso_networks[dest1]
	// If the network exists
	if !ok1 {
		tools.Log(s.Id, dest1+" network does not exist!")
		return
	}
	network2, ok2 := s.Bdso_networks[dest2]
	// If the network exists
	if !ok2 {
		tools.Log(s.Id, dest2+" network does not exist!")
		return
	}

	N1 := len(s.Bdso_networks[dest1])
	F1 := (N1 - 1) / 3
	N2 := len(s.Bdso_networks[dest2])
	F2 := (N2 - 1) / 3

	s.Message_counter++
	m1 := strconv.Itoa(s.Message_counter) + "." + r1
	sendToServers(network1, []string{ADD, m1}, 2*F1+1)
	s.Message_counter++
	m2 := strconv.Itoa(s.Message_counter) + "." + r2
	sendToServers(network2, []string{ADD, m2}, 2*F2+1)
	// WAIT FOR F+1 RESPONSES
	replies1 := make(map[string]bool)
	replies2 := make(map[string]bool)
	tools.Log(s.Id, "Waiting for f+1 ADD_RESPONSE messages on {"+r1+"}")
	tools.Log(s.Id, "Waiting for f+1 ADD_RESPONSE messages on {"+r2+"}")
	for len(replies1) < F1+1 && len(replies2) < F2+1 {
		sockets, _ := s.Poller.Poll(-1)
		for _, socket := range sockets {
			sock := socket.Socket
			msg, _ := sock.RecvMessage(0)
			tools.Log(s.Id, "["+strings.Join(msg, " ")+"]")
			tools.Log(s.Id, "Expected "+r1)
			if msg[1] == ADD_RESPONSE && msg[2] == s.Id+"."+m1 {
				replies1[msg[0]] = true
			}
			if msg[1] == ADD_RESPONSE && msg[2] == s.Id+"."+m2 {
				replies2[msg[0]] = true
			}
			// tools.Log(s.Id, strconv.Itoa(len(replies1))+"/"+strconv.Itoa(F1)+" "+r1)
			// tools.Log(s.Id, strconv.Itoa(len(replies1))+"/"+strconv.Itoa(F2)+" "+r2)
		}
	}
	tools.Log(s.Id, "Record {"+r1+"} appended at "+dest1)
	tools.Log(s.Id, "Record {"+r2+"} appended at "+dest2)
}

func sendToServers(m map[string]*zmq4.Socket, message []string, amount int) {
	sockets := make([]*zmq4.Socket, 0)
	for _, v := range m {
		sockets = append(sockets, v)
	}
	rand.Shuffle(len(sockets), func(i, j int) {
		sockets[i], sockets[j] = sockets[j], sockets[i]
	})
	for _, s := range sockets {
		s.SendMessage(message)
	}
}
