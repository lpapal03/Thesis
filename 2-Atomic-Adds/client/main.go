package main

import (
	"BFT-Distributed-G-Set/client"
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/modules"
	"BFT-Distributed-G-Set/tools"
	"os"
	"strconv"
)

func main() {
	tools.ResetLogFile()
	hosts_filename := "/users/loukis/Thesis/BFT-Distributed-G-Set/hosts"

	servers := config.GetHosts(hosts_filename, "servers")

	config.N = len(servers)
	config.F = (config.N - 1) / 3

	c := client.CreateClient(servers)

	if len(os.Args) < 2 {
		modules.StartInteractive(c)
	} else {
		behaviour := os.Args[1]
		switch behaviour {
		case "interactive":
			modules.StartInteractive(c)
		case "automated":
			if len(os.Args) > 2 {
				request_count, err := strconv.Atoi(os.Args[2])
				if err != nil {
					panic("Invalid arguments")
				}
				modules.StartAutomated(c, request_count)
			} else {
				modules.StartAutomated(c, 20)
			}
		default:
			panic("Invalid arguments")
		}
	}
}
