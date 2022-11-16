// Client

package main

import (
	"frontend/config"
	"frontend/gset"
	"frontend/tools"

	zmq "github.com/pebbe/zmq4"
)

func client_task(id string, servers []config.Server) {

	// Declare context, poller, router sockets of servers, message counter
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	var server_sockets []*zmq.Socket
	message_counter := 0

	// Connect client dealer sockets to all servers
	for i := 0; i < len(servers); i++ {
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(id)
		target := "tcp://" + servers[i].Host + servers[i].Port
		s.Connect(target)
		tools.Log(id, "Established connection with "+target)
		server_sockets = append(server_sockets, s)
		poller.Add(server_sockets[i], zmq.POLLIN)
	}

	// gset.Add(id, server_sockets, &message_counter, poller, "Hello world")

	gset.Get(id, server_sockets, &message_counter, poller)
	// messaging.TargetedMessage([]string{messaging.ADD, "Hello"}, server_sockets[0])

	// messaging.SimpleBroadcast([]string{messaging.GET}, server_sockets)

}

func main() {

	LOCAL := true
	var servers []config.Server
	if LOCAL {
		servers = config.Servers_LOCAL
	} else {
		servers = config.Servers
	}

	go client_task("c1", servers)
	// go client_task("c", servers)
	// go client_task("c1", servers)

	for {
	}

}
