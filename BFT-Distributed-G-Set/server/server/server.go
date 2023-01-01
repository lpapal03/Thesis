package server

import (
	"backend/config"
	"backend/gset"
	"backend/tools"

	zmq "github.com/pebbe/zmq4"
)

type Server struct {
	Zctx           *zmq.Context
	Peers          []*zmq.Socket
	Receive_socket zmq.Socket
	Host           string
	Port           string
	Id             string
	Gset           map[string]string
	BRB            map[string]bool
}

func Create(node config.Node, peers []config.Node) Server {
	id := node.Host + node.Port
	zctx, _ := zmq.NewContext()
	server_sockets := make([]*zmq.Socket, 0)
	my_gset := gset.Create()
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
	}

	return Server{
		Zctx:           zctx,
		Peers:          server_sockets,
		Receive_socket: *receive_socket,
		Host:           node.Host,
		Port:           node.Port,
		Id:             id,
		Gset:           my_gset,
		BRB:            brb,
	}
}
