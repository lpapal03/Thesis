package start

import (
	"bufio"
	"fmt"
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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

	config.CreateScenario("NORMAL", "LOCAL")
	servers := config.SERVERS
	client := client.Create(id, servers)

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

func StartAutomated(client_count, request_count int) {
	var wg sync.WaitGroup
	wg.Add(client_count)
	for i := 0; i < client_count; i++ {
		id := "c" + strconv.Itoa(i)
		go func(id string) {
			defer wg.Done()
			fmt.Println("ID set to '" + id + "'")
			config.CreateScenario("NORMAL", "LOCAL")
			servers := config.SERVERS
			client := client.Create(id, servers)

			time.Sleep(time.Second * 1)
			for r := 0; r < request_count; r++ {
				messaging.Add(client, id+"-ADD-"+strconv.Itoa(r))
				time.Sleep(time.Millisecond * 500)
				messaging.Get(client)
			}
			return
		}(id)
	}
	wg.Wait()
}
