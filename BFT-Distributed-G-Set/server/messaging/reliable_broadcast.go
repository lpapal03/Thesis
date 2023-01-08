package messaging

import (
	"BFT-Distributed-G-Set/server"
	"strings"
)

// Leader, the one who initializes the module
func ReliableBroadcast(leader server.Server, message Message) {
	// send init to everyone
	content := append([]string{message.Sender}, message.Content...)
	tag := BRACHA_BROADCAST_INIT
	v := CreateMessageString(tag, content)
	sendToAll(leader, v)
}

// Called from every server receiving RB messages
func HandleReliableBroadcast(receiver server.Server, v Message) bool {

	// if my vote and my echo are both false, return false

	// peer_echo_key := v.Content[0] + "-" + v.Content[1] + "-" + v.Sender + "-echo"
	// peer_vote_key := v.Content[0] + "-" + v.Content[1] + "-" + v.Sender + "-vote"
	// my_echo_key := v.Content[0] + "-" + v.Content[1] + "-" + receiver.Hostname + "-echo"
	// my_vote_key := v.Content[0] + "-" + v.Content[1] + "-" + receiver.Hostname + "-vote"
	// my_init_key := v.Content[0] + "-" + v.Content[1] + "-" + receiver.Hostname + "-init"
	// bare_key := v.Content[0] + "-" + v.Content[1]

	// // add message in message pot and count
	// if v.Tag == BRACHA_BROADCAST_ECHO {
	// 	receiver.BRB[peer_echo_key] = true
	// }
	// if v.Tag == BRACHA_BROADCAST_VOTE {
	// 	receiver.BRB[peer_vote_key] = true
	// }

	// echo_count, vote_count := countMessages(receiver.BRB, bare_key)

	// // tools.Log(receiver.Id, "Echo: "+strconv.Itoa(echo_count))
	// // tools.Log(receiver.Id, "Vote: "+strconv.Itoa(vote_count))

	// // on receiving <v> from leader
	// if v.Tag == BRACHA_BROADCAST_INIT {
	// 	if receiver.BRB[my_init_key] == false {
	// 		receiver.BRB[my_echo_key] = true
	// 		receiver.BRB[my_vote_key] = true
	// 		v := CreateMessageString(BRACHA_BROADCAST_ECHO, v.Content)
	// 		sendToAll(receiver, v)
	// 		receiver.BRB[my_echo_key] = false
	// 		receiver.BRB[my_init_key] = true
	// 	}
	// }

	// // on receiving <echo, v> from n-f distinct parties:
	// if v.Tag == BRACHA_BROADCAST_ECHO && echo_count >= config.N-config.F {
	// 	if receiver.BRB[my_vote_key] == true {
	// 		v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
	// 		sendToAll(receiver, v)
	// 	}
	// 	receiver.BRB[my_vote_key] = false
	// }

	// // on receiving <echo, v> from f+1 distinct parties:
	// if v.Tag == BRACHA_BROADCAST_ECHO && vote_count >= config.F+1 {
	// 	if receiver.BRB[my_vote_key] == true {
	// 		v := CreateMessageString(BRACHA_BROADCAST_VOTE, v.Content)
	// 		sendToAll(receiver, v)
	// 	}
	// 	receiver.BRB[my_vote_key] = false
	// }

	// // on receiving <vote, v> from n-f distinct parties:
	// if v.Tag == BRACHA_BROADCAST_VOTE && vote_count >= config.N-config.F {
	// 	tools.Log(receiver.Id, "Delivered "+strings.Join(v.Content, " "))
	// 	// clean map (not important, just saves memory)
	// 	potCleanUp(receiver.BRB, bare_key)
	// 	return true
	// }

	// // for k, v := range receiver.BRB {
	// // 	fmt.Println(receiver.Id, k, v)
	// // }

	return false

}

func sendToAll(receiver server.Server, message []string) {
	for _, peer_socket := range receiver.Peers {
		peer_socket.SendMessage(message)
	}
}

func potCleanUp(pot map[string]bool, bare_key string) {

	for k := range pot {
		if strings.Contains(k, bare_key) {
			delete(pot, k)
		}
	}
}
