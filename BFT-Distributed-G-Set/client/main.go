package main

import (
	"BFT-Distributed-G-Set/client"
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/modules"
	"BFT-Distributed-G-Set/tools"
	"os"
)

func main() {
	tools.ResetLogFile()
	filename := "/users/loukis/Thesis/BFT-Distributed-G-Set/hosts"

	servers := config.GetHosts(filename, "servers")

	config.N = len(servers)
	config.F = (config.N - 1) / 3

	c := client.CreateClient(servers)

	behaviour := os.Args[1]
	switch behaviour {
	case "interactive":
		modules.StartInteractive(c)
	case "automated":
		modules.StartAutomated(c)
	default:
		panic("Invalid argument")
	}
}
