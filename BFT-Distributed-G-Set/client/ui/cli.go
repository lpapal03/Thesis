package ui

import (
	"fmt"
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"strings"
)

func client_task(id string, servers []config.Node) {

	client := client.Create(id, servers)

	// messaging.Add(client, msg)
	// messaging.TargetedAdd(client, *client.Servers[0], "Hello")

	messaging.Add(client, "Hello")
	messaging.Add(client, "World")
	messaging.Add(client, "How")
	messaging.Add(client, "Are")
	messaging.Add(client, "You")

	// time.Sleep(time.Millisecond * 500)

	messaging.Get(client)

}

func Start_CLI() {

	// ask about scenario
	// then begin infinite loop

	id := "c1"
	config.CreateScenario("NORMAL", "LOCAL")
	servers := config.SERVERS
	client := client.Create(id, servers)

	var command string
	var record string
	for {
		fmt.Print("Type 'g' for GET, 'a' for ADD or 'e' for EXIT: ")
		fmt.Scanln(&command)
		command = strings.ToLower(command)
		if command == "e" {
			return
		}
		if command == "g" {
			messaging.Get(client)
		}
		if command == "a" {
			fmt.Print("Record to append: ")
			fmt.Scanln(&record)
			messaging.Add(client, record)
		}
	}

	// config.CreateScenario("NORMAL", "LOCAL")
	// servers := config.SERVERS

	// go client_task("c1", servers)

	// // Infinite loop in main thread
	// for {
	// }
}
