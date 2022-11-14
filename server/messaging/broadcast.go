package messaging

import zmq "github.com/pebbe/zmq4"

// Implement reliable broadcast
// Here, all of the sending and receiving will happen
func ReliableBroadcast(message []string, servers []*zmq.Socket, poller *zmq.Poller) {
	for i := 0; i < len(servers); i++ {
		servers[i].SendMessage(message)
	}
}
