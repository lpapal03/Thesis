package client

import (
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/tools"

	zmq "github.com/pebbe/zmq4"
)

type Client struct {
	Id              string
	Zctx            *zmq.Context
	Poller          *zmq.Poller
	Message_counter int
	Servers         map[string]*zmq.Socket
}

func CreateClient(id string, servers []config.Node) *Client {
	// Declare context, poller, router sockets of servers, message counter
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	server_sockets := make(map[string]*zmq.Socket)

	// Connect client dealer sockets to all servers
	for i := 0; i < len(servers); i++ {
		s, err := zctx.NewSocket(zmq.DEALER)
		if err != nil {
			panic(err)
		}
		s.SetIdentity(id)
		target := "tcp://" + servers[i].Host + ":" + servers[i].Port
		s.Connect(target)
		tools.Log(id, "Established connection with "+target)
		server_sockets[servers[i].Host+":"+servers[i].Port] = s
		poller.Add(s, zmq.POLLIN)
	}

	return &Client{
		Id:              id,
		Zctx:            zctx,
		Poller:          poller,
		Message_counter: 0,
		Servers:         server_sockets}

}
