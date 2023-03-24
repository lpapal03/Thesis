package modules

import (
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/messaging"
	"2-Atomic-Adds/server"
	"2-Atomic-Adds/tools"
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

func StartMute(servers []config.Node, default_port, num_threads int) {
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
			tools.Log(s.Id, "Started with MUTE behaviour")
			for {
				msg, err := s.Receive_socket.RecvMessage(0)
				if err != nil {
					panic(err)
				}
				// messaging.HandleMessage(s, msg)
				tools.Log(s.Id, "Received {"+strings.Join(msg, " ")+"}, no action")
			}
		}(i)
	}
}

func StartMalicious(servers []config.Node, default_port, num_threads int) {
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
			tools.Log(s.Id, "Started with MALICIOUS behaviour")
			for {
				msg, err := s.Receive_socket.RecvMessage(0)
				if err != nil {
					panic(err)
				}
				// messaging.HandleMessage(s, msg)
				messaging.HandleMessageByzantine(s, msg, "MALICIOUS")
			}
		}(i)
	}
}

func StartHalfAndHalf(servers []config.Node, default_port, num_threads int) {
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
			tools.Log(s.Id, "Started with HALF_AND_HALF behaviour")
			for {
				msg, err := s.Receive_socket.RecvMessage(0)
				if err != nil {
					panic(err)
				}
				messaging.HandleMessageByzantine(s, msg, "HALF_AND_HALF")
			}
		}(i)
	}
}
