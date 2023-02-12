package client

import (
	"BFT-Distributed-G-Set/config"
	"BFT-Distributed-G-Set/tools"
	"os"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

type Client struct {
	Hostname        string
	Zctx            *zmq.Context
	Poller          *zmq.Poller
	Message_counter int
	Servers         map[string]*zmq.Socket
}

func CreateClient(servers []config.Node) *Client {
	// Declare context, poller, router sockets of servers, message counter
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	server_sockets := make(map[string]*zmq.Socket)

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = strings.Split(hostname, ".")[0]

	// Connect client dealer sockets to all servers
	for i := 0; i < len(servers); i++ {
		s, err := zctx.NewSocket(zmq.DEALER)
		if err != nil {
			panic(err)
		}
		s.SetIdentity(hostname)
		target := "tcp://" + servers[i].Host + ":" + servers[i].Port
		s.Connect(target)
		tools.Log(hostname, "Established connection with "+target)
		server_sockets[servers[i].Host+":"+servers[i].Port] = s
		poller.Add(s, zmq.POLLIN)
	}

	return &Client{
		Hostname:        hostname,
		Zctx:            zctx,
		Poller:          poller,
		Message_counter: 0,
		Servers:         server_sockets}

}
