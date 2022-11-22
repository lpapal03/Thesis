package messaging

import (
	"backend/config"
	"backend/server"
	"backend/tools"
	"fmt"
	"strconv"
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
	leader.BRB_state.My_echo_state[my_states_key] = false
	leader.BRB_state.My_vote_state[my_states_key] = true

}

// Called from every server receiving RB messages
func HandleReliableBroadcast(receiver server.Server, v Message) {

	fmt.Println(receiver.Id, " ", receiver.BRB_state.My_vote_state)

	my_states_key := strings.Join(v.Content, " ")                  // c1 Hello
	pier_pots_key := strings.Join(v.Content, " ") + " " + v.Sender // c1 Hello localhost:1000

	// add message in message pot and count
	if v.Tag == BRACHA_BROADCAST_ECHO {
		receiver.BRB_state.Pier_echo_pot[pier_pots_key] = true
	}
	if v.Tag == BRACHA_BROADCAST_VOTE {
		receiver.BRB_state.Pier_vote_pot[pier_pots_key] = true
	}
	echo_count := countMessages(receiver.BRB_state.Pier_echo_pot, my_states_key)
	vote_count := countMessages(receiver.BRB_state.Pier_vote_pot, my_states_key)

	// on receiving <v> from leader
	if v.Tag == BRACHA_BROADCAST_INIT {
		receiver.BRB_state.My_echo_state[my_states_key] = true
		receiver.BRB_state.My_vote_state[my_states_key] = true
		v := CreateMessageString(BRACHA_BROADCAST_ECHO, v.Content)
		for _, pier_socket := range receiver.Piers {
			pier_socket.SendMessage(v)
		}
		receiver.BRB_state.My_echo_state[my_states_key] = false
	}

	// on receiving <echo, v> from n-f distinct parties:
	if echo_count >= config.N-config.F {
		if receiver.BRB_state.My_vote_state[my_states_key] == true {
			v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
			for _, pier_socket := range receiver.Piers {
				pier_socket.SendMessage(v)
			}
		}
		receiver.BRB_state.My_vote_state[my_states_key] = false
	}

	// on receiving <echo, v> from f+1 distinct parties:
	if vote_count >= config.F+1 {
		if receiver.BRB_state.My_vote_state[my_states_key] == true {
			v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
			for _, pier_socket := range receiver.Piers {
				pier_socket.SendMessage(v)
			}
		}
		receiver.BRB_state.My_vote_state[my_states_key] = false
	}

	if vote_count >= config.N-config.F {
		tools.Log(receiver.Id, "DELIVERED DELIVERED DELIVERED "+strings.Join(v.Content, " "))
	}

	tools.Log(receiver.Id, "Echo: "+strconv.Itoa(echo_count)+"/"+strconv.Itoa(config.N-config.F))
	tools.Log(receiver.Id, "Vote: "+strconv.Itoa(vote_count)+"/"+strconv.Itoa(config.N-config.F))

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
func sendToAll(receiver server.Server, message Message) {
	for _, pier_socket := range receiver.Piers {
		pier_socket.SendMessage(message)
	}
}
