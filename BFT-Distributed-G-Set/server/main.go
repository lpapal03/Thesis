package main

import (
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/server"
	"fmt"
)

func main() {

	filename := "/users/loukis/Thesis/BFT-Distributed-G-Set/hosts"

	servers := config.GetHosts(filename, "servers")

	config.N = len(servers)
	config.F = (config.N - 1) / 3

	server := server.CreateServer(servers)

	fmt.Println(server.Peers)

	for {
		msg, _ := server.Receive_socket.RecvMessage(0)
		fmt.Println(msg)
		server.Receive_socket.SendMessage([]string{msg[0], "Received " + msg[1]})
	}

}
