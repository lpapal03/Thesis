package messaging

import zmq "github.com/pebbe/zmq4"

func SimpleBroadcast(message []string, servers []*zmq.Socket) {
	for i := 0; i < len(servers); i++ {
		servers[i].SendMessage(message)
	}
}

func TargetedMessage(message []string, server *zmq.Socket) {
	server.SendMessage(message)
}
