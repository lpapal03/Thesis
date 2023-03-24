package modules

import (
	"backend/messaging"
	"backend/server"
	"backend/tools"
	"fmt"
)

func Normal_listener_task(s *server.Server) {
	tools.Log(s.Id, "Behaviour: Normal")
	for {
		message, err := s.Receive_socket.RecvMessage(0)
		if err != nil {
			fmt.Println(err)
			return
		}
		messaging.HandleMessage(s, message)
	}
}

func Mute_listener_task(s *server.Server) {
	tools.Log(s.Id, "Behaviour: Mute")
	for {
		msg, err := s.Receive_socket.RecvMessage(0)
		message, err := messaging.ParseMessageString(msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		tools.Log(s.Id, "Received "+message.Tag+" from "+message.Sender+", no action")
	}
}

// Send the same wrong message to everyone
func Malicious_listener_task(s *server.Server) {
	tools.Log(s.Id, "Behaviour: Malicious")
	for {
		message, err := s.Receive_socket.RecvMessage(0)
		if err != nil {
			fmt.Println(err)
			return
		}
		messaging.HandleMessageByzantine(s, message, "MALICIOUS")
	}
}

// Send the same wrong message to half of the servers, and another to the rest
func Half_and_Half_listener_task(s *server.Server) {
	tools.Log(s.Id, "Behaviour: Half and half")
	for {
		message, err := s.Receive_socket.RecvMessage(0)
		if err != nil {
			fmt.Println(err)
			return
		}
		messaging.HandleMessageByzantine(s, message, "HALF_AND_HALF")
	}
}
