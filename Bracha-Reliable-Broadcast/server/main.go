package main

import (
	"backend/config"
	"backend/modules"
	"fmt"
)

func main() {

	config.CreateScenario("NORMAL", "LOCAL")
	servers := config.SERVERS

	// Start all listener servers
	fmt.Println("Initializing...")
	for i := 1; i < config.N; i++ {
		go modules.Listener_task(servers[i], servers)
	}
	modules.Leader_task(servers[0], servers)
	for {
	}

}
