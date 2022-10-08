// Server

package main

import (
	"fmt"
	"log"
	"strconv"

	zmq "github.com/pebbe/zmq4"
)

func get(g_set map[string]string) string {
	v := ""
	for _, value := range g_set {
		v += value + "\n"
	}
	return v
}

func server_task() {
	my_id := 101010
	g_set := make(map[string]string)
	g_set["record1"] = "dog"
	g_set["record2"] = "blue"
	g_set["record3"] = "cat"
	g_set["record4"] = "red"

	zctx, _ := zmq.NewContext()
	inbound_socket, _ := zctx.NewSocket(zmq.ROUTER)
	inbound_socket.Bind("tcp://*:5555")

	// oubound_sockets =
	// for server in all servers:
	// 	add dealer to it
	log.Println("Server is up!")

	for {
		msg, _ := inbound_socket.RecvMessage(0)
		fmt.Println(msg[0])
		fmt.Println(msg[1])
		fmt.Println(msg[2])

		if msg[2] == "get" {
			v := get(g_set)
			response := []string{msg[0], strconv.Itoa(my_id), v}
			inbound_socket.SendMessage(response)
		} else if msg[2] == "append" {
			fmt.Println("Append not implemented yet")
		}
	}
}

func main() {
	go server_task()
	for {
	}
}
