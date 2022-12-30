package main

import (
	"backend/config"
	"backend/modules"
)

func main() {

	config.CreateScenario("NORMAL", "LOCAL")
	servers := config.SERVERS

	// Start all servers with vision of all other servers
	for i := 0; i < config.N; i++ {
		go modules.Server_Task_Normal(servers[i], servers)
	}
	// Infinite loop in main thread
	for {
	}

}
