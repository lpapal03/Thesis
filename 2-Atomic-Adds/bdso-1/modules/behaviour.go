package modules

import (
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/messaging"
	"BFT-Distributed-G-Set/server"
	"BFT-Distributed-G-Set/tools"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StartNormal(servers []config.Node, default_port, num_threads int) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = strings.Split(hostname, ".")[0]
	for i := default_port; i < default_port+num_threads; i++ {
		go func(my_port int) {
			p := strconv.Itoa(my_port)
			me := config.Node{Host: hostname, Port: p}
			s := server.CreateServer(me, servers)
			tools.Log(s.Id, "Started with NORMAL behaviour")
			for {
				msg, err := s.Receive_socket.RecvMessage(0)
				if err != nil {
					panic(err)
				}
				messaging.HandleMessage(s, msg)
			}
		}(i)
	}

}

func StartMute(s *server.Server) {
	tools.Log(s.Id, "Started with MUTE behaviour")
	for {
		msg, _ := s.Receive_socket.RecvMessage(0)
		fmt.Println("Received", msg, " No action")
	}
}

func StartMalicious(s *server.Server) {
	tools.Log(s.Id, "Started with MALICIOUS behaviour")

}
