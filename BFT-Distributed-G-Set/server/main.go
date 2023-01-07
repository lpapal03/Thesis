package main

import (
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/server"
	"BFT-Distributed-G-Set/tools"
	"os"
)

func main() {

	tools.ResetLogFile()
	hosts_filename := "/users/loukis/Thesis/BFT-Distributed-G-Set/hosts"

	peers := config.GetHosts(hosts_filename, "servers")

	config.N = len(peers) + 1
	config.F = (config.N - 1) / 3

	s := server.CreateServer(peers)

	behaviour := os.Args[1]
	switch behaviour {
	case "normal":
		server.StartNormal(s)
	case "mute":
		server.StartNormal(s)
	case "malicious":
		server.StartNormal(s)
	default:
		panic("Invalid argument")
	}

}
