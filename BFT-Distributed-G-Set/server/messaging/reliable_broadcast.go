package messaging

import (
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/server"
	"BFT-Distributed-G-Set/tools"
	"fmt"
	"strconv"
)

// Leader, the one who initializes the module
func ReliableBroadcast(leader server.Server, message Message) {
	content := append([]string{message.Sender}, message.Content...)
	tag := BRACHA_BROADCAST_INIT
	v := CreateMessageString(tag, content)
	// leader with input v
	sendToAll(leader, v)
}

// Called from every server receiving RB messages
func HandleReliableBroadcast(receiver server.Server, v Message) bool {

	my_key := v.Content[1]
	peers_key := v.Sender + "." + v.Content[1]

	fmt.Println(my_key, peers_key)

	// Party j (including the leader)
	if v.Tag == BRACHA_BROADCAST_INIT && !receiver.My_init[my_key] {
		receiver.My_echo[my_key] = true
		receiver.My_vote[my_key] = true
	}

	// on receiving <v> from leader:
	if v.Tag == BRACHA_BROADCAST_INIT {
		if receiver.My_echo[my_key] {
			v := CreateMessageString(BRACHA_BROADCAST_ECHO, v.Content)
			sendToAll(receiver, v)
			receiver.My_echo[my_key] = false
		}
	}

	// count messages
	if v.Tag == BRACHA_BROADCAST_ECHO {
		receiver.Peers_echo[peers_key] = true
	}
	if v.Tag == BRACHA_BROADCAST_VOTE {
		receiver.Peers_vote[peers_key] = true
	}
	// count messages
	s := ""
	for k := range receiver.My_echo {
		s += k + "\n"
	}
	tools.Log(receiver.Hostname, s)
	echo_count, vote_count := countMessages(receiver, peers_key)

	// on receiving <echo, v> from n-f distinct parties:
	if v.Tag == BRACHA_BROADCAST_ECHO && echo_count > config.N-config.F {
		if receiver.My_vote[my_key] {
			v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
			sendToAll(receiver, v)
			receiver.My_vote[my_key] = false
		}
	}

	// on receiving <vote, v> from f+1 distinct parties:
	if v.Tag == BRACHA_BROADCAST_VOTE && vote_count > config.F+1 {
		if receiver.My_vote[my_key] {
			v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
			sendToAll(receiver, v)
			receiver.My_vote[my_key] = false
		}
	}

	// on receiving <vote, v> from n-f distinct parties:
	if v.Tag == BRACHA_BROADCAST_VOTE && vote_count > config.N-config.F {
		return true
	}

	tools.Log(receiver.Hostname, "Echo: "+strconv.Itoa(echo_count))
	tools.Log(receiver.Hostname, "Vote: "+strconv.Itoa(vote_count))

	return false

}

func countMessages(s server.Server, key string) (int, int) {
	echo_count := 0
	vote_count := 0
	for k, v := range s.Peers_echo {
		if k == key && v {
			echo_count++
		}
	}
	for k, v := range s.Peers_vote {
		if k == key && v {
			vote_count++
		}
	}
	return echo_count, vote_count
}

func sendToAll(receiver server.Server, message []string) {
	for _, peer_socket := range receiver.Peers {
		peer_socket.SendMessage(message)
	}
}
