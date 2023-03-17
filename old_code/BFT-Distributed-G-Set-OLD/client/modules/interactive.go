package modules

import (
	"BFT-Distributed-G-Set/client"
	"BFT-Distributed-G-Set/messaging"
	"BFT-Distributed-G-Set/tools"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartInteractive(c *client.Client) {
	tools.Log(c.Hostname, "Started interactive session")
	scanner := bufio.NewScanner(os.Stdin)
	var command string
	var record string

	fmt.Print("Type 'g' for GET, 'a' for ADD or 'e' for EXIT\n> ")
	for scanner.Scan() {
		command = strings.ToLower(scanner.Text())
		if command == "e" {
			return
		}
		if command == "g" {
			messaging.Get(c)
		}
		if command == "a" {
			fmt.Print("Record to append > ")
			scanner.Scan()
			record = scanner.Text()
			if isRecordValid(record) {
				messaging.Add(c, record)
			} else {
				fmt.Print("Record cannot be empty or contain only spaces\n")
			}
		}
		if len(command) == 0 {
			fmt.Print("> ")
		} else {
			fmt.Print("Type 'g' for GET, 'a' for ADD or 'e' for EXIT\n> ")
		}
	}
}
