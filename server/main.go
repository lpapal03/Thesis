package main

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
	"backend/tools"
	"strings"
)

func server_task(me config.Node, piers []config.Node) {

	server := server.Create(me, piers)

	// Listen to messages and handle them
	for {
		message, _ := server.Receive_socket.RecvMessage(0)
		messaging.HandleMessage(server, message)
	}
}

func byzantine_server_task(me config.Node, piers []config.Node) {

	server := server.Create(me, piers)
	tools.Log(server.Id, "Malicious server created")

	// Listen to messages and handle them
	for {
		// message, _ := server.Receive_socket.RecvMessage(0)
		// messaging.HandleMessage(server, message)
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
		if strings.Contains(servers[i].Port, "1000") {
			go server_task(servers[i], servers)
		}
		if strings.Contains(servers[i].Port, "2000") {
			go byzantine_server_task(servers[i], servers)
		}

		// turning point
		// change task from server_task to byzantine_server_task
		if strings.Contains(servers[i].Port, "9999") {
			go server_task(servers[i], servers)
		}
	}
	// Infinite loop in main thread to allow the other threads to run
	for {
	}

}
