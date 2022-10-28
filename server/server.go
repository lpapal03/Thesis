// Server

package main

import (
	"fmt"
	"server/config"
	"server/gset"

	zmq "github.com/pebbe/zmq4"
)

func server_task(my_port string, server_ports []string) {
	// Declare context, poller, router sockets of servers
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	var server_sockets []*zmq.Socket

	// Create gset object
	mygset := gset.Create()
	gset.Append(mygset, "A") //DELETE

	// My router socket
	inbound_socket, _ := zctx.NewSocket(zmq.ROUTER)
	inbound_socket.Bind("tcp://*:" + my_port)
	fmt.Printf("Port %s is bound\n", my_port)

	// Connect server dealer sockets to all other servers
	for i := 0; i < len(server_ports); i++ {
		// Connect if not me
		if server_ports[i] == my_port {
			continue
		}
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(my_port)
		s.Connect("tcp://localhost:" + server_ports[i])
		server_sockets = append(server_sockets, s)
		poller.Add(server_sockets[len(server_sockets)-1], zmq.POLLIN)
		// fmt.Printf("Server %s connected to server %s\n", my_port, server_ports[i])
	}

	for {
		msg, _ := inbound_socket.RecvMessage(0)
		fmt.Println(msg)
		response := []string{msg[0], "World", my_port}
		inbound_socket.SendMessage(response)
	}
}

func main() {
	// Start all servers
	for i := 0; i < config.N; i++ {
		go server_task(config.Servers[i], config.Servers)
	}
	// Infinite loop in main thread to allow processes to run
	for {
	}

}
