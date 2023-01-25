package scenarios

import (
	"backend/messaging"
	"backend/server"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func Normal_listener_task(s *server.Server) {
	for {
		message, err := s.Receive_socket.RecvMessage(0)
		if err != nil {
			fmt.Println(zmq.AsErrno(err))
			panic(err)
		}
		messaging.HandleMessage(s, message)
	}
}
