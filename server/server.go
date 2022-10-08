// Server

package main

import (
	"fmt"
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

func server_task(id string, context *zmq.Context) {

	g_set := make(map[string]string)
	g_set["record1"] = "dog"
	g_set["record2"] = "blue"
	g_set["record3"] = "cat"
	g_set["record4"] = "red"

	server_count := 5

	inbound_socket, _ := context.NewSocket(zmq.ROUTER)
	port := "tcp://*:5555"
	inbound_socket.Bind(port)
	fmt.Println("Client facing socket is up in port: ", port)

	oubound_sockets := make([]*zmq.Socket, 0)
	for i := 0; i < server_count; i++ {
		s, _ := context.NewSocket(zmq.DEALER)
		port := "tcp://*:1000" + strconv.Itoa(i)
		s.Bind(port)
		fmt.Println("Bound dealer to port:", port)
		oubound_sockets = append(oubound_sockets, s)
	}
	fmt.Println("Server facing sockets are up!")

	for {
		msg, _ := inbound_socket.RecvMessage(0)
		fmt.Println(msg[0])
		fmt.Println(msg[1])
		fmt.Println(msg[2])

		if msg[2] == "get" {
			v := get(g_set)
			response := []string{msg[0], id, v}
			inbound_socket.SendMessage(response)
		} else if msg[2] == "append" {
			fmt.Println("Append not implemented yet")
		}
	}
}

func main() {
	zctx, _ := zmq.NewContext()
	go server_task("s1", zctx)
	for {
	}
}
