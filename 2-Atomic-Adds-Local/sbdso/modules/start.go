package modules

import (
	"backend/config"
	"backend/server"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func Start(servers []config.Node, scenario string, zctx *zmq.Context, bdso_networks map[string][]config.Node) {
	// all servers act normal
	if scenario == "NORMAL" {
		for i := 0; i < config.N; i++ {
			s := server.CreateServer(servers[i], servers, zctx, bdso_networks)
			go Normal_listener_task(s)
		}
		return
	}
	if scenario == "MUTES" {
		fmt.Println("Scenario not implemented")
		return
	}
	if scenario == "HALF&HALF" {
		fmt.Println("Scenario not implemented")
		return
	}
	if scenario == "MALICIOUS" {
		fmt.Println("Scenario not implemented")
		return
	}
}
