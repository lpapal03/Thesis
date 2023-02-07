package main

import (
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/modules"
	"BFT-Distributed-G-Set/tools"
	"strconv"
)

func main() {
	tools.ResetLogFile()
	wd := "/users/loukis/Thesis/BFT-Distributed-G-Set"

	// hosts are just the machine names
	hosts := config.GetHosts(wd+"/hosts", "servers")
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

	modules.StartNormal(servers, default_port, num_threads)

	// if len(os.Args) < 2 {
	// 	modules.StartNormal(servers, default_port, num_threads)
	// } else {
	// 	behaviour := os.Args[1]
	// 	switch behaviour {
	// 	case "normal":
	// 		modules.StartNormal(default_port, num_threads)
	// 	case "mute":
	// 		modules.StartMute(default_port, num_threads)
	// 	case "malicious":
	// 		modules.StartMute(default_port, num_threads)
	// 	default:
	// 		panic("Invalid argument")
	// 	}
	// }

	select {}

}
