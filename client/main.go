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

	// messaging.Add(client, "Hello")

	msg := ""

	if client.Id == "c1" {
		msg = "Hello aaaa"
	}
	if client.Id == "c2" {
		msg = "World"
	}
	if client.Id == "c3" {
		msg = "How_are_you"
	}

	messaging.TargetedAdd(client, *client.Servers[0], msg)
	time.Sleep(time.Second * 2)
	messaging.TargetedAdd(client, *client.Servers[0], msg)
	g, _ := messaging.GetGset(client)
	fmt.Println(g)
	// for {
	// 	messaging.TargetedAdd(client, *client.Servers[0], "Hello")
	// 	time.Sleep(time.Second * 2)
	// }

	// time.Sleep(time.Second * 3)
	// messaging.TargetedAdd(client, *client.Servers[0], "Hello")
	// messaging.TargetedAdd(client, *client.Servers[1], "Hello22")
	// s, _ := messaging.GetGset(client)
	// fmt.Println(s)
	// messaging.SimpleBroadcast([]string{messaging.GET}, server_sockets)

}

func main() {

	LOCAL := false
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
