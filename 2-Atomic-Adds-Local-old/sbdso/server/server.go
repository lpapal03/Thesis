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
	Receive_socket *zmq.Socket
	Id             string
	Gset           map[string]string
	Port           string
	My_init        map[string]bool
	My_echo        map[string]bool
	My_vote        map[string]bool
	Peers_echo     map[string]bool
	Peers_vote     map[string]bool
	Bdso_networks  map[string]map[string]*zmq.Socket
}

// Now server receives a map of ports with their respective bdso
func CreateServer(node config.Node, peers []config.Node, zctx *zmq.Context, bdso_network map[string][]config.Node) *Server {
	id := node.Host + node.Port
	port := node.Port
	peer_sockets := make(map[string]*zmq.Socket)
	my_gset := gset.Create()
	my_init := make(map[string]bool)
	my_echo := make(map[string]bool)
	my_vote := make(map[string]bool)
	peers_echo := make(map[string]bool)
	peers_vote := make(map[string]bool)
	bdso_net := make(map[string]map[string]*zmq.Socket)
	receive_socket, _ := zctx.NewSocket(zmq.ROUTER)
	receive_socket.Bind("tcp://*:" + node.Port)
	tools.Log(id, "Bound tcp://*:"+node.Port)

	// Connect my dealer sockets to all other servers' router
	for i := 0; i < len(peers); i++ {
		s, err := zctx.NewSocket(zmq.DEALER)
		if err != nil {
			panic(err)
		}
		s.SetIdentity(id)
		s.Connect("tcp://localhost:" + peers[i].Port)
		// append socket to socket list
		peer_sockets["tcp://localhost:"+peers[i].Port] = s
		// tools.Log(id, "Connected to "+"tcp://localhost:"+peers[i].Port)
	}

	for network_name := range bdso_network {
		tools.Log(id, "Starting connection with network: "+network_name)
		bdso_net[network_name] = make(map[string]*zmq.Socket)
		for _, node_id := range bdso_network[network_name] {
			s, err := zctx.NewSocket(zmq.DEALER)
			if err != nil {
				panic(err)
			}
			s.SetIdentity(id)
			s.Connect("tcp://localhost:" + node_id.Port)
			bdso_net[network_name]["tcp://localhost:"+node_id.Port] = s
			tools.Log(id, "Connected to "+"tcp://localhost:"+node_id.Port+" of network "+network_name)
		}
	}

	return &Server{
		Peers:          peer_sockets,
		Receive_socket: receive_socket,
		Id:             id,
		Port:           port,
		Gset:           my_gset,
		My_init:        my_init,
		My_echo:        my_echo,
		My_vote:        my_vote,
		Peers_echo:     peers_echo,
		Peers_vote:     peers_vote,
		Bdso_networks:  bdso_net,
	}
}
