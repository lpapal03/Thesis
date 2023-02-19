package scenarios

import (
	"backend/messaging"
	"backend/server"
	"fmt"
)

func Normal_listener_task(s *server.Server) {
	for {
		message, err := s.Receive_socket.RecvMessage(0)
		if err != nil {
			fmt.Println(s.Receive_socket)
			panic(err)
		}
		messaging.HandleMessage(s, message)
	}
}
