package ui

import (
	"bufio"
	"fmt"
	"frontend/client"
	"frontend/messaging"
	"os"
	"strings"
)

func Start_CLI() {

	scanner := bufio.NewScanner(os.Stdin)
	var id string
	var command string
	var record string

	// fmt.Print("Your ID\n> ")
	// scanner.Scan()
	// id = scanner.Text()
	// fmt.Println("ID set to '" + scanner.Text() + "'")

	// FOR TESTING
	id = "c1"
	fmt.Println("ID set to '" + id + "'")

	config.CreateScenario("NORMAL", "LOCAL")
	servers := config.SERVERS
	client := client.Create(id, servers)

	// var command string
	// var record string
	fmt.Print("Type 'g' for GET, 'a' for ADD or 'e' for EXIT\n> ")
	for scanner.Scan() {
		command = strings.ToLower(scanner.Text())
		if command == "e" {
			return
		}
		if command == "g" {
			messaging.Get(client)
		}
		if command == "a" {
			fmt.Print("Record to append\n> ")
			scanner.Scan()
			record = scanner.Text()
			messaging.Add(client, record)
		}
		if len(command) == 0 {
			fmt.Print("> ")
		} else {
			fmt.Print("Type 'g' for GET, 'a' for ADD or 'e' for EXIT\n> ")
		}
	}
}
