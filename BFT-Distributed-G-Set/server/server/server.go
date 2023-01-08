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
	Receive_socket zmq.Socket
	Hostname       string
	Gset           map[string]string

	Echo_pot map[string]bool
	Vote_pot map[string]bool
}

func CreateServer(peers []string) Server {

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = strings.Split(hostname, ".")[0]

	zctx, _ := zmq.NewContext()
	server_sockets := make(map[string]*zmq.Socket)
	my_gset := gset.Create()

	gset.Append(my_gset, "INIT")

	echo_pot := make(map[string]bool)
	vote_pot := make(map[string]bool)

	receive_socket, _ := zctx.NewSocket(zmq.ROUTER)
	receive_socket.Bind("tcp://*:" + config.DEFAULT_PORT)
	tools.Log(hostname, "Bound tcp://*:"+config.DEFAULT_PORT)

	// Connect my dealer sockets to all other servers' router
	for i := 0; i < len(peers); i++ {
		// Connect if not me
		if peers[i] == hostname {
			continue
		}
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(hostname)
		s.Connect("tcp://" + peers[i])
		tools.Log(hostname, "Connected to "+peers[i])
		// append socket to socket list
		server_sockets[peers[i]] = s
	}

	return Server{
		Zctx:           zctx,
		Peers:          server_sockets,
		Receive_socket: *receive_socket,
		Hostname:       hostname,
		Gset:           my_gset,
		Echo_pot:       echo_pot,
		Vote_pot:       vote_pot,
	}
}
