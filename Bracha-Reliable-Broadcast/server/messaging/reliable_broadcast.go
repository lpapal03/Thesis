package messaging

import (
	"backend/config"
	"backend/server"
	"backend/tools"
	"strings"
)

// Leader, the one who initializes the module
func ReliableBroadcast(leader server.Server, message Message) {

	tools.Log(leader.Id, "Called Reliable broadcast")

	my_echo_key := message.Sender + "-" + message.Content[0] + "-" + leader.Id + "-echo"
	my_vote_key := message.Sender + "-" + message.Content[0] + "-" + leader.Id + "-vote"
	leader.BRB[my_echo_key] = true
	leader.BRB[my_vote_key] = true

	content := append([]string{message.Sender}, message.Content...)

	// send init to everyone
	tag := BRACHA_BROADCAST_INIT
	v := CreateMessageString(tag, content)
	sendToAll(leader, v)

	// send echo to everyone (assume I received INIT from self)
	tag = BRACHA_BROADCAST_ECHO
	v = CreateMessageString(tag, content)
	sendToAll(leader, v)
	leader.BRB[my_echo_key] = true
}

// Called from every server receiving RB messages
func HandleReliableBroadcast(receiver server.Server, v Message) bool {

	peer_echo_key := v.Content[0] + "-" + v.Content[1] + "-" + v.Sender + "-echo"
	peer_vote_key := v.Content[0] + "-" + v.Content[1] + "-" + v.Sender + "-vote"
	my_echo_key := v.Content[0] + "-" + v.Content[1] + "-" + receiver.Id + "-echo"
	my_vote_key := v.Content[0] + "-" + v.Content[1] + "-" + receiver.Id + "-vote"
	count_key := v.Content[0] + "-" + v.Content[1]

	// add message in message pot and count
	if v.Tag == BRACHA_BROADCAST_ECHO {
		receiver.BRB[peer_echo_key] = true
	}
	if v.Tag == BRACHA_BROADCAST_VOTE {
		receiver.BRB[peer_vote_key] = true
	}

	echo_count, vote_count := countMessages(receiver.BRB, count_key)

	// on receiving <v> from leader
	if v.Tag == BRACHA_BROADCAST_INIT {
		receiver.BRB[my_echo_key] = true
		receiver.BRB[my_vote_key] = true
		v := CreateMessageString(BRACHA_BROADCAST_ECHO, v.Content)
		sendToAll(receiver, v)
		receiver.BRB[my_echo_key] = false
	}

	// on receiving <echo, v> from n-f distinct parties:
	if echo_count >= config.N-config.F {
		if receiver.BRB[my_vote_key] == true {
			v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
			sendToAll(receiver, v)
		}
		receiver.BRB[my_vote_key] = false
	}

	// on receiving <echo, v> from f+1 distinct parties:
	if vote_count >= config.F+1 {
		if receiver.BRB[my_vote_key] == true {
			v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
			sendToAll(receiver, v)
		}
		receiver.BRB[my_vote_key] = false
	}

	// on receiving <vote, v> from n-f distinct parties:
	if vote_count >= config.N-config.F {
		// clean map (not important, just saves memory)
		receiver.BRB = make(map[string]bool)
		return true
	}

	return false

}

// count the messages received for a given v
func countMessages(pot map[string]bool, count_key string) (int, int) {
	// start counters from 1, assuming caller is true on echo and vote
	echo_count := 1
	vote_count := 1
	for k, v := range pot {

		if strings.Contains(k, count_key) && strings.Contains(k, "echo") && v {
			echo_count++
		}
		if strings.Contains(k, count_key) && strings.Contains(k, "vote") && v {
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
