package server

import (
	"backend/config"
	"backend/gset"
	"backend/tools"

	zmq "github.com/pebbe/zmq4"
)

type Server struct {
	Zctx           *zmq.Context
	Poller         *zmq.Poller
	Peers          []*zmq.Socket
	Receive_socket zmq.Socket
	Host           string
	Port           string
	Id             string
	Gset           map[string]string
	// Echo_state     map[string]bool
	// Vote_state     map[string]bool
	// Peer_echo_pot  map[string]bool
	// Peer_vote_pot  map[string]bool
	BRB map[string]bool
}

func Create(node config.Node, peers []config.Node) Server {
	id := node.Host + node.Port
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	server_sockets := make([]*zmq.Socket, 0)
	my_gset := gset.Create()
	// echo_state := make(map[string]bool)
	// vote_state := make(map[string]bool)
	// peer_echo_pot := make(map[string]bool)
	// peer_vote_pot := make(map[string]bool)
	brb := make(map[string]bool)
	receive_socket, _ := zctx.NewSocket(zmq.ROUTER)
	receive_socket.Bind("tcp://*:" + node.Port)
	tools.Log(id, "Bound tcp://*:"+node.Port)

	// Connect my dealer sockets to all other servers' router
	for i := 0; i < len(peers); i++ {
		// Connect if not me
		if peers[i] == node {
			continue
		}
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(id)
		s.Connect("tcp://" + peers[i].Host + peers[i].Port)
		// append socket to socket list
		server_sockets = append(server_sockets, s)
		// new socket is the last one
		poller.Add(s, zmq.POLLIN)
	}

	return Server{
		Zctx:           zctx,
		Poller:         poller,
		Peers:          server_sockets,
		Receive_socket: *receive_socket,
		Host:           node.Host,
		Port:           node.Port,
		Id:             id,
		Gset:           my_gset,
		BRB:            brb,
	}
}
