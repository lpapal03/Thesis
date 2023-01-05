package main

import (
	"fmt"
	"frontend/client"
	"os"
	"strings"
)

func main() {

	// fmt.Println("Hello from client")

	data, err := os.ReadFile("/users/loukis/Thesis/BFT-Distributed-G-Set-V2/hosts")
	if err != nil {
		panic(err)
	}
	hosts := strings.Split(strings.ReplaceAll(string(data), "\n\n", "\n"), "\n")
	servers := hosts[:len(hosts)-1]
	for i := 0; i < len(servers); i++ {
		if servers[i] == "[servers]" {
			servers = servers[i+1:]
			break
		}
	}

	client := client.CreateClient(servers)
		for i:=0; i<len(client.Servers); i++{
			client.Servers[i].SendMessage([]string{"Hello"})
		}
	fmt.Println("Sent hello, waiting for reply...")
	for {
		sockets, _ := client.Poller.Poll(-1)
		for _, socket := range sockets {
			s := socket.Socket
			msg, _ := s.RecvMessage(0)
			fmt.Println(msg)
		}
	}

}
