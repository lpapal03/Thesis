package main

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
)

func server_task(me config.Node, peers []config.Node) {

	server := server.Create(me, peers)

	// Listen to messages and handle them
	for {
		message, _ := server.Receive_socket.RecvMessage(0)
		messaging.HandleMessage(server, message)
	}
}

func main() {
	config.CreateScenario("NORMAL", "LOCAL")
	servers := config.SERVERS

	// Start all servers with vision of all other servers
	for i := 0; i < config.N; i++ {
		go server_task(servers[i], servers)

	}
	// Infinite loop in main thread
	for {
	}

}
