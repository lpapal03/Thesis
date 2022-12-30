package modules

import (
	"backend/config"
	"fmt"
)

func Start(servers []config.Node, scenario string) {
	// all servers act normal
	if scenario == "NORMAL" {
		for i := 1; i < config.N; i++ {
			go Normal_listener_task(servers[i], servers)
		}
		Leader_task(servers[0], servers)
		for {
		}

	}
	// f mutes
	if scenario == "MUTES" {
		// 1 leader
		go Leader_task(servers[0], servers)
		fmt.Println("Leader(Normal): ", servers[0])
		// f mute
		for i := 1; i < config.F+1; i++ {
			go Mute_listener_task(servers[i], servers)
			fmt.Println("Mute: ", servers[i])
		}
		// n-f normal
		for i := config.F + 1; i < config.N; i++ {
			go Normal_listener_task(servers[i], servers)
			fmt.Println("Normal: ", servers[i])
		}
		for {
		}
	}
	// f/2 act correctly, the rest act maliciously
	if scenario == "HALF&HALF" {

	}
}
