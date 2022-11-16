// Client

package main

import (
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
)

func client_task(id string, servers []config.Node) {

	client := client.Create(id, servers)

	// gset.Add(id, server_sockets, &message_counter, poller, "Hello world")

	messaging.GetGset(client)
	// messaging.TargetedMessage([]string{messaging.ADD, "Hello"}, server_sockets[0])

	// messaging.SimpleBroadcast([]string{messaging.GET}, server_sockets)

}

func main() {

	LOCAL := true
	var servers []config.Node
	if LOCAL {
		servers = config.Servers_LOCAL
	} else {
		servers = config.Servers
	}

	go client_task("c1", servers)

	for {
	}

}
