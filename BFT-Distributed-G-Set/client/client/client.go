package client

import (
	"frontend/config"
	"frontend/tools"
	"os"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

type Client struct {
	Id              string
	Zctx            zmq.Context
	Poller          zmq.Poller
	Message_counter int
	Servers         []*zmq.Socket
}

func CreateClient(servers []string) Client {
	// Declare context, poller, router sockets of servers, message counter
	zctx, _ := zmq.NewContext()
	poller := zmq.NewPoller()
	var server_sockets []*zmq.Socket

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	id := strings.Split(hostname, ".")[0]

	// Connect client dealer sockets to all servers
	for i := 0; i < len(servers); i++ {
		s, _ := zctx.NewSocket(zmq.DEALER)
		s.SetIdentity(id)
		target := "tcp://" + servers[i] + ":" + config.DEFAULT_PORT
		s.Connect(target)
		tools.Log(id, "Established connection with "+target)
		server_sockets = append(server_sockets, s)
		poller.Add(s, zmq.POLLIN)
	}

	return Client{
		Id:              id,
		Zctx:            *zctx,
		Poller:          *poller,
		Message_counter: 0,
		Servers:         server_sockets}

}
