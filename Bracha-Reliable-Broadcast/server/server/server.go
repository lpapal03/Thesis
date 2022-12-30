package server

import (
	"backend/config"
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
	BRB            map[string]bool
}

func Create(node config.Node, peers []config.Node) Server {
	id := node.Host + node.Port
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	server_sockets := make([]*zmq.Socket, 0)
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
		BRB:            brb,
	}
}
