package modules

import (
	"backend/messaging"
	"backend/server"
)

func Normal_listener_task(s *server.Server) {
	for {
		message, err := s.Receive_socket.RecvMessage(0)
		if err != nil {
			panic(err)
		}
		messaging.HandleMessage(s, message)
	}
}
