// Client

package main

import (
	"fmt"
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
)

func client_task(id string, servers []config.Node) {

	client := client.Create(id, servers)

	// messaging.Add(client, msg)
	messaging.TargetedAdd(client, *client.Servers[0], "1")
	// time.Sleep(time.Second * 2)
	g, _ := messaging.GetGset(client)
	fmt.Println(g)

}

func main() {

	config.CreateScenario("NORMAL", "LOCAL")
	servers := config.SERVERS

	go client_task("c1", servers)

	// Infinite loop in main thread
	for {
	}

}
