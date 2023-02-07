package server

import (
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/gset"
	"BFT-Distributed-G-Set/tools"
	"os"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

type Server struct {
	Zctx           *zmq.Context
	Peers          map[string]*zmq.Socket
	Receive_socket *zmq.Socket
	Hostname       string
	Gset           map[string]string
	My_init        map[string]bool
	My_echo        map[string]bool
	My_vote        map[string]bool
	Peers_echo     map[string]bool
	Peers_vote     map[string]bool
}

func CreateServer(peers []string) *Server {

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = strings.Split(hostname, ".")[0]

	zctx, _ := zmq.NewContext()
	server_sockets := make(map[string]*zmq.Socket)
	my_gset := gset.Create()
	my_init := make(map[string]bool)
	my_echo := make(map[string]bool)
	my_vote := make(map[string]bool)
	peers_echo := make(map[string]bool)
	peers_vote := make(map[string]bool)

	receive_socket, _ := zctx.NewSocket(zmq.ROUTER)
	receive_socket.Bind("tcp://*:" + config.DEFAULT_PORT)
	tools.Log(hostname, "Bound tcp://*:"+config.DEFAULT_PORT)

	// Connect my dealer sockets to all other servers' router
	for i := 0; i < len(peers); i++ {
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(hostname)
		s.Connect("tcp://" + peers[i] + ":" + config.DEFAULT_PORT)
		tools.Log(hostname, "Connected to "+peers[i]+":"+config.DEFAULT_PORT)
		// append socket to socket list
		server_sockets[peers[i]] = s
	}

	return &Server{
		Zctx:           zctx,
		Peers:          server_sockets,
		Receive_socket: receive_socket,
		Hostname:       hostname,
		Gset:           my_gset,
		My_init:        my_init,
		My_echo:        my_echo,
		My_vote:        my_vote,
		Peers_echo:     peers_echo,
		Peers_vote:     peers_vote,
	}
}
