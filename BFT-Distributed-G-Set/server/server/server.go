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
	Peers          []*zmq.Socket
	Receive_socket zmq.Socket
	Hostname       string
	Id             string
	Gset           map[string]string
	BRB            map[string]bool
}

func Create(peers []string) Server {

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = strings.Split(hostname, ".")[0]

	id := hostname + config.DEFAULT_PORT
	zctx, _ := zmq.NewContext()
	server_sockets := make([]*zmq.Socket, 0)
	my_gset := gset.Create()
	brb := make(map[string]bool)
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
		s.SetIdentity(id)
		s.Connect("tcp://" + peers[i])
		tools.Log(hostname, "Connected to "+peers[i])
		// append socket to socket list
		server_sockets = append(server_sockets, s)
	}

	return Server{
		Zctx:           zctx,
		Peers:          server_sockets,
		Receive_socket: *receive_socket,
		Id:             id,
		Hostname:       hostname,
		Gset:           my_gset,
		BRB:            brb,
	}
}
