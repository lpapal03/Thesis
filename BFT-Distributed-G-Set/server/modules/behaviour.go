package modules

import (
	"BFT-Distributed-G-Set/server"
	"BFT-Distributed-G-Set/tools"
	"fmt"
)

func StartNormal(s server.Server) {
	tools.Log(s.Hostname, "Started with NORMAL behaviour")
	for {
		msg, _ := s.Receive_socket.RecvMessage(0)
		fmt.Println(msg)
		s.Receive_socket.SendMessage([]string{msg[0], "Received " + msg[1]})
	}
}

func StartMute(s server.Server) {
	tools.Log(s.Hostname, "Started with MUTE behaviour")
	for {
		msg, _ := s.Receive_socket.RecvMessage(0)
		fmt.Println("Received", msg, " No action")
	}
}

func StartMalicious(s server.Server) {
	tools.Log(s.Hostname, "Started with MALICIOUS behaviour")

}
