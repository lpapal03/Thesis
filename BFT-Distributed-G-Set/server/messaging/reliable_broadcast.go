package messaging

import (
	"backend/config"
	"backend/gset"
	"backend/server"
	"backend/tools"
	"strings"
)

// Leader, the one who initializes the module
func ReliableBroadcast(leader server.Server, message Message) {

	tools.Log(leader.Id, "Called Reliable broadcast module")

	content := append([]string{message.Sender}, message.Content...)

	// send init to everyone
	tag := BRACHA_BROADCAST_INIT
	v := CreateMessageString(tag, content)
	for _, peer_socket := range leader.Peers {
		peer_socket.SendMessage(v)
	}

	// send echo to everyone (assume I received INIT from self)
	tag = BRACHA_BROADCAST_ECHO
	v = CreateMessageString(tag, content)
	for _, peer_socket := range leader.Peers {
		peer_socket.SendMessage(v)
	}

	states_key := message.Content[0]
	leader.Echo_state[states_key] = true
	leader.Vote_state[states_key] = false
}

// Called from every server receiving RB messages
func HandleReliableBroadcast(receiver server.Server, v Message) bool {

	// if exists stop. No true effect, just
	// improves performance
	if gset.Exists(receiver.Gset, v.Content[1]) {
		return false
	}

	states_key := v.Content[1]                     // c1 Hello
	peer_pots_key := v.Content[1] + " " + v.Sender // c1 Hello localhost:1000

	// add message in message pot and count
	if v.Tag == BRACHA_BROADCAST_ECHO {
		receiver.Peer_echo_pot[peer_pots_key] = true
	}
	if v.Tag == BRACHA_BROADCAST_VOTE {
		receiver.Peer_vote_pot[peer_pots_key] = true
	}

	echo_count := countMessages(receiver.Peer_echo_pot, states_key)
	vote_count := countMessages(receiver.Peer_vote_pot, states_key)

	// on receiving <v> from leader
	if v.Tag == BRACHA_BROADCAST_INIT {
		receiver.Echo_state[states_key] = true
		receiver.Vote_state[states_key] = true
		v := CreateMessageString(BRACHA_BROADCAST_ECHO, v.Content)
		sendToAll(receiver, v)
		receiver.Echo_state[states_key] = false
	}

	// on receiving <echo, v> from n-f distinct parties:
	if echo_count >= config.N-config.F {
		if receiver.Vote_state[states_key] == true {
			v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
			sendToAll(receiver, v)
		}
		receiver.Vote_state[states_key] = false
	}

	// on receiving <echo, v> from f+1 distinct parties:
	if vote_count >= config.F+1 {
		if receiver.Vote_state[states_key] == true {
			v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
			sendToAll(receiver, v)
		}
		receiver.Vote_state[states_key] = false
	}

	// on receiving <vote, v> from n-f distinct parties:
	if vote_count >= config.N-config.F {
		tools.Log(receiver.Id, "Delivered "+strings.Join(v.Content, " "))
		return true
	}

	// tools.Log(receiver.Id, "Echo: "+strconv.Itoa(echo_count)+"/"+strconv.Itoa(config.N-config.F))
	// tools.Log(receiver.Id, "Vote: "+strconv.Itoa(vote_count)+"/"+strconv.Itoa(config.N-config.F))
	return false

}

// count the messages received for a given v
func countMessages(pot map[string]bool, identifier string) int {
	// assume I received echo and vote from self
	count := 1
	for key := range pot {
		if strings.Contains(key, identifier) {
			count++
		}
	}
	return count
}

func sendToAll(receiver server.Server, message []string) {
	for _, peer_socket := range receiver.Peers {
		peer_socket.SendMessage(message)
	}
}

// can be used to destroy unused objects
func brb_state_cleanup(server server.Server, identifier string) {

	delete(server.Echo_state, identifier)
	delete(server.Vote_state, identifier)

	for k := range server.Peer_echo_pot {
		if strings.Contains(k, identifier) {
			delete(server.Peer_echo_pot, k)
		}
	}

	for k := range server.Peer_vote_pot {
		if strings.Contains(k, identifier) {
			delete(server.Peer_vote_pot, k)
		}
	}
}
