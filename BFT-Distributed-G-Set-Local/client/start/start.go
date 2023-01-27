package start

import (
	"bufio"
	"fmt"
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"frontend/tools"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func waitRandomly(min, max int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max - min)
	time.Sleep(time.Duration(min+r) * time.Millisecond)
}

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

func automated_client_task(id string, req_count int) {
	fmt.Println("ID set to '" + id + "'")
	config.CreateScenario("NORMAL", "LOCAL")
	servers := config.SERVERS
	client := client.CreateClient(id, servers)

	time.Sleep(time.Second * 1)
	for r := 0; r < req_count; r++ {
		messaging.Add(client, id+"-"+strconv.Itoa(r))
		// waitRandomly(1000, 2000)
		messaging.Get(client)
		// waitRandomly(1000, 2000)
	}
	tools.Log(id, "Done")
}

func StartAutomated(client_count, request_count int) {
	var wg sync.WaitGroup
	wg.Add(client_count)
	for i := 0; i < client_count; i++ {
		id := "c" + strconv.Itoa(i)
		go func(id string) {
			fmt.Println("ID set to '" + id + "'")
			config.CreateScenario("NORMAL", "LOCAL")
			servers := config.SERVERS
			client := client.CreateClient(id, servers)

			time.Sleep(time.Second * 1)
			for r := 0; r < request_count; r++ {
				messaging.Add(client, id+"-"+strconv.Itoa(r))
				waitRandomly(1000, 2000)
				messaging.Get(client)
				waitRandomly(1000, 2000)
			}
			tools.Log(id, "Done")
			wg.Done()
		}(id)
	}
	wg.Wait()
}
