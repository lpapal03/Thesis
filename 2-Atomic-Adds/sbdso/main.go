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
	wd := "/users/loukis/Thesis/2-Atomic-Adds"

	// hosts are just the machine names
	sbdso := config.GetHosts(wd+"/hosts", "servers")
	bdso_1 := config.GetHosts(wd+"/hosts", "bdso-1")
	bdso_2 := config.GetHosts(wd+"/hosts", "bdso-2")
	default_port, num_threads := config.GetPortAndThreads(wd + "/config")

	servers := make([]config.Node, 0)
	for _, h := range hosts {
		for p := default_port; p < default_port+num_threads; p++ {
			p_num := strconv.Itoa(p)
			servers = append(servers, config.Node{Host: h, Port: p_num})
		}
	}

	config.N = len(servers)
	config.F = (config.N - 1) / 3

	if len(os.Args) < 2 {
		modules.StartNormal(servers, default_port, num_threads)
	} else {
		behaviour := os.Args[1]
		switch behaviour {
		case "normal":
			modules.StartNormal(servers, default_port, num_threads)
		case "mute":
			modules.StartMute(servers, default_port, num_threads)
		case "malicious":
			modules.StartMute(servers, default_port, num_threads)
		default:
			panic("Invalid argument")
		}
	}

	select {}

}
