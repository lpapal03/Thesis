package modules

import (
	"BFT-Distributed-G-Set/client"
	"BFT-Distributed-G-Set/messaging"
	"BFT-Distributed-G-Set/tools"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func isRecordValid(r string) bool {
	return !(r == "" || strings.TrimSpace(r) == "")
}

func StartInteractive(c client.Client) {
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
				fmt.Print("Record cannot be empty or contain only spaces")
			}
		}
		if len(command) == 0 {
			fmt.Print("> ")
		} else {
			fmt.Print("Type 'g' for GET, 'a' for ADD or 'e' for EXIT\n> ")
		}
	}
}

func StartAutomated(c client.Client) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func(c client.Client) {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			messaging.Get(c)
			rand.Seed(time.Now().UnixNano())
			t := rand.Intn(4)
			time.Sleep(time.Duration(t) * time.Second)
		}
	}(c)

	go func(c client.Client) {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(6)
			b := make([]byte, n)
			for i := range b {
				b[i] = 'a' + byte(rand.Intn(26))
			}
			s := string(b)
			if isRecordValid(s) {
				messaging.Add(c, s)
			}
			rand.Seed(time.Now().UnixNano())
			t := rand.Intn(4)
			time.Sleep(time.Duration(t) * time.Second)
		}
	}(c)

	wg.Wait()

}
