package main

import (
	"BFT-Distributed-G-Set/client"
	"BFT-Distributed-G-Set/config"
	"fmt"
)

func main() {

	filename := "/users/loukis/Thesis/BFT-Distributed-G-Set/hosts"

	servers := config.GetHosts(filename, "servers")

	config.N = len(servers)
	config.F = (config.N - 1) / 3

	client := client.CreateClient(servers)

	fmt.Println(client.Servers)
	// ui.Start_CLI(client)
}
