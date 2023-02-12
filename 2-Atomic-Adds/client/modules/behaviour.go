package modules

import (
	"2-Atomic-Adds/client"
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/messaging"
	"2-Atomic-Adds/tools"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func isMessageValid(msg string) bool {
	if msg == "" {
		return false
	}
	if strings.Contains(msg, " ") {
		return false
	}
	if strings.Contains(msg, ".") {
		return false
	}
	if strings.Contains(msg, "{") {
		return false
	}
	if strings.Contains(msg, "}") {
		return false
	}
	if strings.Contains(msg, ";") {
		return false
	}
	return true
}

func isAtomicMessageValid(msg string) bool {
	if msg == "" {
		return false
	}
	if strings.Contains(msg, " ") {
		return false
	}
	if strings.Contains(msg, ".") {
		return false
	}
	if strings.Contains(msg, "{") {
		return false
	}
	if strings.Contains(msg, "}") {
		return false
	}
	parts := strings.Split(msg, ";")
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		if len(p) < 1 {
			return false
		}
	}
	if !config.NetworkExists(strings.Split(msg, ";")[1]) {
		return false
	}
	return true
}

func waitRandomly(min, max int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max - min)
	time.Sleep(time.Duration(min+r) * time.Millisecond)
}

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

	fmt.Print("Type 'g' for GET, 'a' for ADD, 'at' for ATOMIC-ADD or 'e' for EXIT\n> ")
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
			if isMessageValid(record) {
				messaging.Add(client, record)
			} else {
				fmt.Println("Invalid message")
			}
		}
		if command == "at" {
			fmt.Println("Format of atomic records: peer_id;destination;your_message;peer_message")
			fmt.Print("Record to append atomically > ")
			scanner.Scan()
			record = scanner.Text()
			if network_name != "sbdso" {
				fmt.Println("Network does not allow atomic operations")
			} else if isAtomicMessageValid(record) {
				messaging.AddAtomic(client, record)
			} else {
				fmt.Println("Invalid message")
			}
		}
		if len(command) == 0 {
			fmt.Print("> ")
		} else {
			fmt.Print("Type 'g' for GET, 'a' for ADD, 'at' for ATOMIC-ADD or 'e' for EXIT\n> ")
		}
	}
}

func StartAutomated(zctx *zmq.Context, client_count, request_count int, network_name string) {
	servers := config.Initialize(network_name)
	var wg sync.WaitGroup
	wg.Add(client_count)
	for i := 0; i < client_count; i++ {
		id := "c" + strconv.Itoa(i)
		go func(id string) {
			tools.Log(id, "Id set")
			config.Initialize(network_name)
			client := client.CreateClient(id, servers, zctx)

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
