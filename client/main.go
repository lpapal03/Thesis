// Client

package main

import (
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
)

func client_task(id string, servers []config.Node) {

	client := client.Create(id, servers)

	// messaging.Add(client, "Hello")

	messaging.TargetedAdd(client, *client.Servers[0], "Hello")

	// messaging.GetGset(client)
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
