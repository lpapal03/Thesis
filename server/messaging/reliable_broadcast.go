package messaging

import (
	"backend/config"
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
	v := CreateMessage(BRACHA_BROADCAST_INIT, append([]string{message.Sender}, message.Content...))
	for _, pier_sokcet := range leader.Piers {
		pier_sokcet.SendMessage(v)
	}
	tools.Log(leader.Id, "Sent message {"+(strings.Join(v, " "))+"} to all piers")

}

// Called from every server receiving RB messages
func HandleReliableBroadcast(receiver server.Server, v Message) bool {

	// maps piers to answer types
	// if message received in poller is equal to v, add it to maps
	init := make(map[string]bool)
	echo := make(map[string]bool)
	ready := make(map[string]bool)

	a, b, c := countMessages(init, echo, ready)
	fmt.Println(a, b, c)

	fmt.Println("HandleReliableBroadcast for ", v)

	// step 1
	// for !isStepDone(1, init, echo, ready) {
	// 	sockets, _ := receiver.Poller.Poll(-1)
	// 	for _, socket := range sockets {

	// 	}
	// }
	sendToAll(receiver, v, BRACHA_BROADCAST_ECHO)

	// if all is good, send true, meaning message
	// has been correctly processed
	return true

}

func sendToAll(receiver server.Server, message Message, new_tag string) {
	msg := CreateMessage(new_tag, append([]string{message.Sender}, message.Content...))
	for _, pier_sokcet := range receiver.Piers {
		pier_sokcet.SendMessage(msg)
	}
}

func countMessages(init, echo, ready map[string]bool) (int, int, int) {
	init_count := 0
	echo_count := 0
	ready_count := 0

	for _, v := range init {
		if v {
			init_count++
		}
	}
	for _, v := range echo {
		if v {
			echo_count++
		}
	}
	for _, v := range ready {
		if v {
			ready_count++
		}
	}

	return init_count, echo_count, ready_count
}

func isStepDone(step int, init, echo, ready map[string]bool) bool {

	init_count, echo_count, ready_count := countMessages(init, echo, ready)

	if step == 1 && (init_count > 0 || echo_count >= (config.N+config.F)/2 || ready_count >= config.F+1) {
		return true
	}
	if step == 2 && (echo_count >= (config.N+config.F)/2 || ready_count >= config.F+1) {
		return true
	}
	if step == 3 && (ready_count >= 2*config.F+1) {
		return true
	}

	return false
}
