package modules

import (
	"bufio"
	"fmt"
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"os"
	"strings"
)

func StartInteractive() {

	scanner := bufio.NewScanner(os.Stdin)
	var id string
	var command string
	var record string

	fmt.Print("Your ID\n> ")
	scanner.Scan()
	id = scanner.Text()
	fmt.Println("ID set to '" + id + "'\n")

	config.Initialize()
	servers := config.SERVERS
	client := client.CreateClient(id, servers)

	fmt.Print("Type 'g' for GET, 'a' for ADD or 'e' for EXIT\n> ")
	for scanner.Scan() {
		command = strings.ToLower(scanner.Text())
		if command == "e" {
			os.Exit(0)
		}
		if command == "g" {
			messaging.Get(client)
		}
		if command == "a" {
			fmt.Print("Record to append > ")
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
