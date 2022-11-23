package server

import (
	"backend/config"
	"backend/gset"
	"backend/tools"

	zmq "github.com/pebbe/zmq4"
)

type brb_state struct {
	My_init_state    map[string]bool
	My_echo_state    map[string]bool
	My_vote_state    map[string]bool
	My_deliver_state map[string]bool
	Pier_echo_pot    map[string]bool
	Pier_vote_pot    map[string]bool
}

type Server struct {
	Zctx           *zmq.Context
	Poller         *zmq.Poller
	Piers          []*zmq.Socket
	Receive_socket zmq.Socket
	Host           string
	Port           string
	Id             string
	Gset           map[string]string
	BRB_state      brb_state
}

func Create(node config.Node, piers []config.Node) Server {
	id := node.Host + node.Port
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	server_sockets := make([]*zmq.Socket, 0)
	my_gset := gset.Create()
	my_init_state := make(map[string]bool)
	my_echo_state := make(map[string]bool)
	my_vote_state := make(map[string]bool)
	my_deliver_state := make(map[string]bool)
	pier_echo_pot := make(map[string]bool)
	pier_vote_pot := make(map[string]bool)
	gset.Append(my_gset, "INIT")
	receive_socket, _ := zctx.NewSocket(zmq.ROUTER)
	receive_socket.Bind("tcp://*:" + node.Port)
	tools.Log(id, "Bound tcp://*:"+node.Port)

	// Connect my dealer sockets to all other servers' router
	for i := 0; i < len(piers); i++ {
		// Connect if not me
		if piers[i] == node {
			continue
		}
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(id)
		s.Connect("tcp://" + piers[i].Host + piers[i].Port)
		// append socket to socket list
		server_sockets = append(server_sockets, s)
		// new socket is the last one
		poller.Add(s, zmq.POLLIN)
	}

	return Server{
		Zctx:           zctx,
		Poller:         poller,
		Piers:          server_sockets,
		Receive_socket: *receive_socket,
		Host:           node.Host,
		Port:           node.Port,
		Id:             id,
		Gset:           my_gset,
		BRB_state: brb_state{
			My_init_state:    my_init_state,
			My_echo_state:    my_echo_state,
			My_vote_state:    my_vote_state,
			My_deliver_state: my_deliver_state,
			Pier_echo_pot:    pier_echo_pot,
			Pier_vote_pot:    pier_vote_pot}}
}
