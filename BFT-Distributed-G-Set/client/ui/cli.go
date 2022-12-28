package ui

import (
	"fmt"
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"strings"
)

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
}
