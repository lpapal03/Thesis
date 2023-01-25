package server

import (
	"backend/config"
	"backend/gset"
	"backend/tools"

	zmq "github.com/pebbe/zmq4"
)

type Server struct {
	Zctx           *zmq.Context
	Peers          map[string]*zmq.Socket
	Receive_socket zmq.Socket
	Id             string
	Gset           map[string]string
	Port           string

	My_init    map[string]bool
	My_echo    map[string]bool
	My_vote    map[string]bool
	Peers_echo map[string]bool
	Peers_vote map[string]bool
}

func CreateServer(node config.Node, peers []config.Node, zctx *zmq.Context) *Server {
	id := node.Host + node.Port
	port := node.Port
	server_sockets := make(map[string]*zmq.Socket)
	my_gset := gset.Create()
	my_init := make(map[string]bool)
	my_echo := make(map[string]bool)
	my_vote := make(map[string]bool)
	peers_echo := make(map[string]bool)
	peers_vote := make(map[string]bool)
	receive_socket, _ := zctx.NewSocket(zmq.ROUTER)
	receive_socket.Bind("tcp://*:" + node.Port)
	tools.Log(id, "Bound tcp://*:"+node.Port)

	// Connect my dealer sockets to all other servers' router
	for i := 0; i < len(peers); i++ {
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(id)
		s.Connect("tcp://localhost:" + peers[i].Port)
		// append socket to socket list
		server_sockets["tcp://localhost:"+peers[i].Port] = s
	}

	return &Server{
		Peers:          server_sockets,
		Receive_socket: *receive_socket,
		Id:             id,
		Port:           port,
		Gset:           my_gset,
		My_init:        my_init,
		My_echo:        my_echo,
		My_vote:        my_vote,
		Peers_echo:     peers_echo,
		Peers_vote:     peers_vote,
	}
}
