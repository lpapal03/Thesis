package main

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
)

func server_task(me config.Node, piers []config.Node) {

	server := server.Create(me, piers)

	// Listen to messages and handle them
	for {
		message, _ := server.Receive_socket.RecvMessage(0)
		go messaging.HandleMessage(server, messaging.ParseMessage(message))
	}
}

func main() {
	LOCAL := true
	var servers []config.Node
	if LOCAL {
		servers = config.Servers_LOCAL
	} else {
		servers = config.Servers
	}

	// Start all servers with vision of all other servers
	for i := 0; i < config.N; i++ {
		go server_task(servers[i], servers)
	}
	// Infinite loop in main thread to allow the other threads to run
	for {
	}

}
