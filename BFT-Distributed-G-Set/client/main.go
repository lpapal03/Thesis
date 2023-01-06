package main

import (
	"BFT-Distributed-G-Set/client"
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/ui"
	"os"
	"strings"
)

func main() {

	data, err := os.ReadFile("/users/loukis/Thesis/BFT-Distributed-G-Set/hosts")
	if err != nil {
		panic(err)
	}
	hosts := strings.Split(strings.ReplaceAll(string(data), "\n\n", "\n"), "\n")
	servers := hosts[:len(hosts)-1]
	for i := 0; i < len(servers); i++ {
		if servers[i] == "[servers]" {
			servers = servers[i+1:]
			break
		}
	}

	config.N = len(servers)
	config.F = (config.N - 1) / 3

	client := client.CreateClient(servers)

	ui.Start_CLI(client)
}
