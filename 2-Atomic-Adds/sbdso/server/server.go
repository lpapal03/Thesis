package server

import (
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/gset"
	"2-Atomic-Adds/tools"

	zmq "github.com/pebbe/zmq4"
)

type Server struct {
	Zctx           *zmq.Context
	Id             string
	Peers          map[string]*zmq.Socket
	Receive_socket *zmq.Socket
	Poller         *zmq.Poller
	Host           string
	Port           string
	Gset           map[string]string
	My_init        map[string]bool
	My_echo        map[string]bool
	My_vote        map[string]bool
	Peers_echo     map[string]bool
	Peers_vote     map[string]bool
	Bdso_networks  map[string]map[string]*zmq.Socket
}

func CreateServer(me config.Node, peers []config.Node, bdso_networks map[string][]config.Node) *Server {
	zctx, _ := zmq.NewContext()
	server_sockets := make(map[string]*zmq.Socket)
	my_gset := gset.Create()
	my_init := make(map[string]bool)
	my_echo := make(map[string]bool)
	my_vote := make(map[string]bool)
	peers_echo := make(map[string]bool)
	peers_vote := make(map[string]bool)
	poller := zmq.NewPoller()
	bdso_net := make(map[string]map[string]*zmq.Socket)
	my_id := me.Host + ":" + me.Port
	receive_socket, err := zctx.NewSocket(zmq.ROUTER)
	if err != nil {
		panic(err)
	}
	receive_socket.Bind("tcp://*:" + me.Port)
	tools.Log(my_id, "Bound tcp://*:"+me.Port)

	// Connect my dealer sockets to all other servers' router
	for i := 0; i < len(peers); i++ {
		s, err := zctx.NewSocket(zmq.DEALER)
		if err != nil {
			panic(err)
		}
		s.SetIdentity(my_id)
		s.Connect("tcp://" + peers[i].Host + ":" + peers[i].Port)
		tools.Log(my_id, "Connected to "+peers[i].Host+":"+peers[i].Port+" (peer)")
		// append socket to socket list
		server_sockets[peers[i].Host+":"+peers[i].Port] = s
	}

	for network_name := range bdso_networks {
		tools.Log(my_id, "Starting connection with network: "+network_name)
		bdso_net[network_name] = make(map[string]*zmq.Socket)
		for _, node_id := range bdso_networks[network_name] {
			s, err := zctx.NewSocket(zmq.DEALER)
			if err != nil {
				panic(err)
			}
			s.SetIdentity(my_id)
			s.Connect("tcp://" + node_id.Host + ":" + node_id.Port)
			poller.Add(s, zmq.POLLIN)
			bdso_net[network_name]["tcp://"+node_id.Host+":"+node_id.Port] = s
			tools.Log(my_id, "Connected to "+"tcp://"+node_id.Host+":"+node_id.Port+" of network "+network_name)
		}
	}

	return &Server{
		Id:             my_id,
		Peers:          server_sockets,
		Receive_socket: receive_socket,
		Poller:         poller,
		Host:           me.Host,
		Port:           me.Port,
		Gset:           my_gset,
		My_init:        my_init,
		My_echo:        my_echo,
		My_vote:        my_vote,
		Peers_echo:     peers_echo,
		Peers_vote:     peers_vote,
		Bdso_networks:  bdso_net,
	}
}
