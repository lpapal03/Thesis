package server

import (
	"backend/config"
	"backend/gset"
	"backend/tools"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

type Server struct {
	Zctx           *zmq.Context
	Poller         *zmq.Poller
	Piers          []*zmq.Socket
	Receive_socket zmq.Socket
	Host           string
	Port           string
	Id             string
	Gset           map[string]string
	BroadcastState map[string]BroadcastValues
}

// For every Broadcast request, we are going to have 2
// lists with echo and ready values for each peer server.
// When new broadcast request is invoked, broadcast state
// has an entry appended with a request identifier plus the
// echo and ready values. When request is resolved, entry is
// deleted.
type BroadcastValues struct {
	Echo          map[string]bool
	Ready         map[string]bool
	BroadcastStep int
}

func Create(node config.Node, piers []config.Node) Server {
	id := node.Host + node.Port
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	server_sockets := make([]*zmq.Socket, 0)
	my_gset := gset.Create()
	broadcast_state := make(map[string]BroadcastValues)
	gset.Append(my_gset, "Helo")
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
		poller.Add(server_sockets[len(server_sockets)-1], zmq.POLLIN)
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
		BroadcastState: broadcast_state}
}

// I am receiver
func CreateBroadcastState(receiver Server, message_content []string) {

	broadcast_values := BroadcastValues{Echo: make(map[string]bool), Ready: make(map[string]bool), BroadcastStep: 0}
	m := strings.Join(message_content, " ")
	receiver.BroadcastState[m] = broadcast_values
	tools.Log(receiver.Id, "Created broadcast state for message {"+m+"}")
}

func ResetBroadCastState() {

}
