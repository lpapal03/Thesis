package main

import (
	"backend/server"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	// the only thing i know is what i have to do
	// and the servers in the network
	hostname, err := os.Hostname()

	if err != nil {
		panic(err)
	}
	hostname = strings.Split(hostname, ".")[0]

	data, err := os.ReadFile("/users/loukis/Thesis/BFT-Distributed-G-Set-V2/server/hosts")
	if err != nil {
		panic(err)
	}
	hosts := strings.Split(strings.ReplaceAll(string(data), "\n\n", "\n"), "\n")
	hosts = hosts[:len(hosts)-1]
	for i := 0; i < len(hosts); i++ {
		if hosts[i] == "[servers]" {
			hosts = hosts[i+1:]
			break
		}
	}

	server.Create(hostname, hosts)

	time.Sleep(time.Second * 60)
	_, e := os.Create("GeeksforGeeks.txt")
	if e != nil {
		log.Fatal(e)
	}

}
