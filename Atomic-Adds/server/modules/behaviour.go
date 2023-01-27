package modules

import (
	"BFT-Distributed-G-Set/messaging"
	"BFT-Distributed-G-Set/server"
	"BFT-Distributed-G-Set/tools"
	"fmt"
)

func StartNormal(s server.Server) {
	tools.Log(s.Hostname, "Started with NORMAL behaviour")
	for {
		msg, err := s.Receive_socket.RecvMessage(0)
		if err != nil {
			tools.Log(s.Hostname, err.Error())
			panic(err)
		}
		messaging.HandleMessage(s, msg)
		// s.Receive_socket.SendMessage([]string{msg[0], "Received " + msg[1]})
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
