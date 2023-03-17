package main

import (
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/modules"
	"2-Atomic-Adds/tools"
	"os"
	"strconv"
)

func main() {
	tools.ResetLogFile()
	wd := "/users/loukis/Thesis/BFT-Distributed-G-Set-Remote"

	servers := config.GetHosts(wd+"/hosts", "servers")
	default_port, num_threads := config.GetPortAndThreads(wd + "/config")

	sevrer_nodes := make([]config.Node, 0)

	for _, h := range servers {
		for p := default_port; p < default_port+num_threads; p++ {
			p_num := strconv.Itoa(p)
			sevrer_nodes = append(sevrer_nodes, config.Node{Host: h, Port: p_num})
		}
	}

	config.N = len(sevrer_nodes)
	config.F = (config.N - 1) / 3

	if len(os.Args) < 2 {
		modules.StartNormal(sevrer_nodes, default_port, num_threads)
	} else {
		behaviour := os.Args[1]
		switch behaviour {
		case "normal":
			modules.StartNormal(sevrer_nodes, default_port, num_threads)
		case "mute":
			modules.StartMute(sevrer_nodes, default_port, num_threads)
		case "malicious":
			modules.StartMute(sevrer_nodes, default_port, num_threads)
		default:
			panic("Invalid argument")
		}
	}

	select {}

}
