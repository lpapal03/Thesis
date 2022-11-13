package messaging

import zmq "github.com/pebbe/zmq4"

func Broadcast(message []string, servers []*zmq.Socket) {
	for i := 0; i < len(servers); i++ {
		servers[i].SendMessage(message)
	}
}
