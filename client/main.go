// Client

package main

import (
	"fmt"
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"time"
)

func client_task(id string, servers []config.Node) {

	client := client.Create(id, servers)

	msg := "Hello"

	// messaging.Add(client, msg)
	messaging.TargetedAdd(client, *client.Servers[0], msg)
	time.Sleep(time.Second * 2)
	g, _ := messaging.GetGset(client)
	fmt.Println(g)

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
	// go client_task("c2", servers)
	// go client_task("c3", servers)

	for {
	}

}
