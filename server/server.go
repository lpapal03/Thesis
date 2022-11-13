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
	}

	// Listen to messages
	for {
		msg, _ := inbound_socket.RecvMessage(0)

		sender_id := msg[0]
		message_type := msg[1]
		tools.Log(me.Host+me.Port, message_type+" from "+sender_id)

		if message_type == messaging.GET {
			gset.HandleGet(sender_id, me.Host+me.Port, *inbound_socket, mygset)

		}
		if message_type == messaging.ADD {
			gset.HandleAdd(me.Host+me.Port, mygset, msg[2], server_sockets)
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
