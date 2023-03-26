package modules

import (
	"2-Atomic-Adds/client"
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/messaging"
	"2-Atomic-Adds/tools"
	"bufio"
	"fmt"
	"os"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func StartInteractive(zctx *zmq.Context, network_name string) {
	servers := config.Initialize(network_name)

	scanner := bufio.NewScanner(os.Stdin)
	var id string
	var command string
	var record string

	fmt.Print("Your ID\n> ")
	scanner.Scan()
	id = scanner.Text()
	for !isMessageValid(id) {
		fmt.Print("Invalid ID, try again\n> ")
		scanner.Scan()
		id = scanner.Text()
	}
	fmt.Println("ID set to '" + id + "'\n")

	client := client.CreateClient(id, servers, zctx)

	fmt.Print("Type 'g' for GET, 'a' for ADD or 'e' for EXIT\n> ")
	for scanner.Scan() {
		command = strings.ToLower(scanner.Text())
		if command == "e" {
			os.Exit(0)
		}
		if command == "g" {
			r, _ := messaging.Get(client)
			tools.Log(client.Id, r)
		}
		if command == "a" {
			fmt.Print("Record to append > ")
			scanner.Scan()
			record = scanner.Text()
			if isMessageValid(record) {
				messaging.Add(client, record)
			} else {
				fmt.Println("Invalid message")
			}
		}
		if len(command) == 0 {
			fmt.Print("> ")
		} else {
			fmt.Print("Type 'g' for GET, 'a' for ADD or 'e' for EXIT\n> ")
		}
	}
}
