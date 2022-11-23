package messaging

import (
	"backend/config"
	"backend/server"
	"backend/tools"
	"strings"
)

func ReliableBroadcast(leader server.Server, message Message) {

	tools.Log(leader.Id, "Called Reliable broadcast module")

	v := CreateMessageString(BRACHA_BROADCAST_INIT, append([]string{message.Sender}, message.Content...))
	for _, pier_socket := range leader.Piers {
		pier_socket.SendMessage(v)
	}
	v[0] = BRACHA_BROADCAST_ECHO
	for _, pier_socket := range leader.Piers {
		pier_socket.SendMessage(v)
	}

	my_states_key := message.Sender + " " + strings.Join(message.Content, " ")
	brb_state_cleanup(leader, my_states_key)
	leader.BRB_state.My_echo_state[my_states_key] = false
	leader.BRB_state.My_vote_state[my_states_key] = true
}

// Called from every server receiving RB messages
func HandleReliableBroadcast(receiver server.Server, v Message) {

	my_states_key := strings.Join(v.Content, " ")                  // c1 Hello
	pier_pots_key := strings.Join(v.Content, " ") + " " + v.Sender // c1 Hello localhost:1000

	// add message in message pot and count
	if isMessageValid(receiver, my_states_key) {
		if v.Tag == BRACHA_BROADCAST_ECHO {
			receiver.BRB_state.Pier_echo_pot[pier_pots_key] = true
		}
		if v.Tag == BRACHA_BROADCAST_VOTE {
			receiver.BRB_state.Pier_vote_pot[pier_pots_key] = true
		}
	}
	echo_count := countMessages(receiver.BRB_state.Pier_echo_pot, my_states_key)
	vote_count := countMessages(receiver.BRB_state.Pier_vote_pot, my_states_key)

	// on receiving <v> from leader
	if v.Tag == BRACHA_BROADCAST_INIT {
		brb_state_cleanup(receiver, my_states_key)
		receiver.BRB_state.My_init_state[my_states_key] = true
		receiver.BRB_state.My_echo_state[my_states_key] = true
		receiver.BRB_state.My_vote_state[my_states_key] = true
		v := CreateMessageString(BRACHA_BROADCAST_ECHO, v.Content)
		sendToAll(receiver, v)
		receiver.BRB_state.My_echo_state[my_states_key] = false
	}

	// on receiving <echo, v> from n-f distinct parties:
	if echo_count >= config.N-config.F {
		if receiver.BRB_state.My_vote_state[my_states_key] == true {
			v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
			sendToAll(receiver, v)
		}
		receiver.BRB_state.My_vote_state[my_states_key] = false
	}

	// on receiving <echo, v> from f+1 distinct parties:
	if vote_count >= config.F+1 {
		if receiver.BRB_state.My_vote_state[my_states_key] == true {
			v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
			sendToAll(receiver, v)
		}
		receiver.BRB_state.My_vote_state[my_states_key] = false
	}

	if vote_count >= config.N-config.F {
		tools.Log(receiver.Id, "Delivered "+strings.Join(v.Content, " "))
		receiver.BRB_state.My_deliver_state[my_states_key] = true
		handleAddInternal(receiver, v)
		brb_state_cleanup(receiver, my_states_key)
	}

	// tools.Log(receiver.Id, "Echo: "+strconv.Itoa(echo_count)+"/"+strconv.Itoa(config.N-config.F))
	// tools.Log(receiver.Id, "Vote: "+strconv.Itoa(vote_count)+"/"+strconv.Itoa(config.N-config.F))

}

func brb_state_cleanup(server server.Server, identifier string) {

	delete(server.BRB_state.My_init_state, identifier)
	delete(server.BRB_state.My_echo_state, identifier)
	delete(server.BRB_state.My_vote_state, identifier)
	delete(server.BRB_state.My_deliver_state, identifier)

	for k := range server.BRB_state.Pier_echo_pot {
		if strings.Contains(k, identifier) {
			delete(server.BRB_state.Pier_echo_pot, k)
		}
	}

	for k := range server.BRB_state.Pier_vote_pot {
		if strings.Contains(k, identifier) {
			delete(server.BRB_state.Pier_vote_pot, k)
		}
	}

}

func isMessageValid(receiver server.Server, identifier string) bool {
	valid := false
	for key := range receiver.BRB_state.My_init_state {
		if strings.Contains(key, identifier) {
			valid = true
		}
	}
	return valid
}

// count the messages received for a given v
// returns init_count, echo_count, ready_count
func countMessages(pot map[string]bool, identifier string) int {
	// example:
	// identifier = c1 hello
	count := 0
	for key := range pot {
		if strings.Contains(key, identifier) {
			count++
		}
	}
	return count
}

// CHANGEEEE
func sendToAll(receiver server.Server, message []string) {
	for _, pier_socket := range receiver.Piers {
		pier_socket.SendMessage(message)
	}
}
