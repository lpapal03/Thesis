package start

import (
	"bufio"
	"fmt"
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	client := client.Create(id, servers)

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

func StartAutomated(client_count, request_count int) {

	config.CreateScenario("NORMAL", "LOCAL")
	servers := config.SERVERS
	client := client.Create("c", servers)
	for i := 0; i < request_count; i++ {
		messaging.Add(client, client.Id+"-"+strconv.Itoa(i))
		// waitRandomly(1000, 2000)
		messaging.Get(client)
	}

	// var wg sync.WaitGroup
	// wg.Add(client_count)
	// for i := 0; i < client_count; i++ {
	// 	id := "c" + strconv.Itoa(i)
	// 	go func(id string) {
	// 		fmt.Println("ID set to '" + id + "'")
	// 		config.CreateScenario("NORMAL", "LOCAL")
	// 		servers := config.SERVERS
	// 		client := client.Create(id, servers)

	// 		time.Sleep(time.Second * 1)
	// 		for r := 0; r < request_count; r++ {
	// 			messaging.Add(client, id+"-"+strconv.Itoa(r))
	// 			// waitRandomly(1000, 2000)
	// 			messaging.Get(client)
	// 			// waitRandomly(1000, 2000)
	// 		}
	// 		tools.Log(id, "Done")
	// 		wg.Done()
	// 	}(id)
	// }
	// wg.Wait()
}
