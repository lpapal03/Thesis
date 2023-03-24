package modules

import (
	"backend/config"
	"backend/server"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func Start(servers []config.Node, scenario string, zctx *zmq.Context) {
	switch scenario {
	case "NORMAL":
		for i := 0; i < config.N; i++ {
			s := server.CreateServer(servers[i], servers, zctx)
			go Normal_listener_task(s)
		}
		select {}
	case "MUTE":
		for i := 0; i < config.N; i++ {
			s := server.CreateServer(servers[i], servers, zctx)
			if i < config.F {
				go Mute_listener_task(s)
			} else {
				go Normal_listener_task(s)
			}
		}
		select {}
	case "MALICIOUS":
		for i := 0; i < config.N; i++ {
			s := server.CreateServer(servers[i], servers, zctx)
			if i < config.F {
				go Malicious_listener_task(s)
			} else {
				go Normal_listener_task(s)
			}
		}
		select {}
	case "HALF_AND_HALF":
		for i := 0; i < config.N; i++ {
			s := server.CreateServer(servers[i], servers, zctx)
			if i < config.F {
				go Half_and_Half_listener_task(s)
			} else {
				go Normal_listener_task(s)
			}
		}
		select {}
	default:
		fmt.Println("Scenario not supported")
	}

}
