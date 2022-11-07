// Server

package main

import (
	"server/config"
	"server/gset"
	"server/messaging"
	"server/tools"

	zmq "github.com/pebbe/zmq4"
)

func server_task(me config.Server, servers []config.Server) {
	// Declare context, poller, router sockets of servers
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	var server_sockets []*zmq.Socket

	// Create gset object
	mygset := gset.Create()
	gset.Append(mygset, "A")
	gset.Append(mygset, "B")
	gset.Append(mygset, "C")

	// My router socket
	inbound_socket, _ := zctx.NewSocket(zmq.ROUTER)
	inbound_socket.Bind("tcp://*:" + me.Port)
	tools.Log(me.Host+me.Port, "Bound tcp://*:"+me.Port)

	// Connect server dealer sockets to all other servers
	for i := 0; i < len(servers); i++ {
		// Connect if not me
		if servers[i] == me {
			continue
		}
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(me.Host + me.Port)
		s.Connect("tcp://" + servers[i].Host + servers[i].Port)
		server_sockets = append(server_sockets, s)
		poller.Add(server_sockets[len(server_sockets)-1], zmq.POLLIN)
		// fmt.Printf("Server %s connected to server %s\n", my_port, server_ports[i])
	}

	// Listen to messages
	for {
		msg, _ := inbound_socket.RecvMessage(0)
		tools.Log(me.Host+me.Port, msg[1]+" from "+msg[0])
		if msg[1] == messaging.GET {
			// msg[0] = sender_id
			response := []string{msg[0], me.Host + me.Port, messaging.GET_RESPONSE, gset.GsetToString(mygset)}
			inbound_socket.SendMessage(response)
			tools.Log(me.Host+me.Port, messaging.GET_RESPONSE+" to "+msg[0])
		}
	}
}

func main() {
	LOCAL := true
	var servers []config.Server
	if LOCAL {
		servers = config.Servers_LOCAL
	} else {
		servers = config.Servers
	}

	// Start all servers
	for i := 0; i < config.N; i++ {
		go server_task(servers[i], servers)
	}
	// Infinite loop in main thread to allow the other threads to run
	for {
	}

}
