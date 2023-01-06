package ui

import (
	"BFT-Distributed-G-Set/client"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printReply(c client.Client) {
	sockets, _ := c.Poller.Poll(-1)
	for _, socket := range sockets {
		s := socket.Socket
		msg, _ := s.RecvMessage(0)
		fmt.Println(msg)
	}
}

func Start_CLI(c client.Client) {

	scanner := bufio.NewScanner(os.Stdin)
	var command string

	fmt.Print("Send to servers:\n> ")
	for scanner.Scan() {
		command = strings.ToLower(scanner.Text())
		if len(command) == 0 {
			fmt.Print("> ")
		} else {
			for i := 0; i < len(c.Servers); i++ {
				c.Servers[i].SendMessage([]string{command})
			}
			fmt.Printf("Sent %s, waiting for reply...\n", command)
			sockets, _ := c.Poller.Poll(-1)
			for _, socket := range sockets {
				s := socket.Socket
				msg, _ := s.RecvMessage(0)
				fmt.Println(msg)
			}
		}
		fmt.Print("Send to servers:\n> ")
	}
}
