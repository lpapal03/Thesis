package main

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
	"backend/tools"
	"bufio"
	"fmt"
	"os"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	servers := config.SetServers()

	zctx, _ := zmq.NewContext()
	// start servers
	for i := 0; i < config.N; i++ {
		go func(listener config.Node, peers []config.Node, zctx *zmq.Context) {
			server := server.CreateServer(listener, peers, zctx)
			for {
				message, err := server.Receive_socket.RecvMessage(0)
				if err != nil {
					panic(err)
				}
				messaging.HandleMessage(server, message)
			}
		}(servers[i], servers, zctx)
	}

	s, _ := zctx.NewSocket(zmq.DEALER)
	s.SetIdentity("DEFAULT_CLIENT")
	target := "tcp://" + servers[0].Host + servers[0].Port
	s.Connect(target)
	tools.Log("DEFAULT_CLIENT", "Established connection with "+target)

	time.Sleep(time.Second * 1)
	fmt.Print("\nValue to broadcast: ")
	scanner.Scan()
	value := scanner.Text()

	s.SendMessage([]string{messaging.BRACHA_BROADCAST, value})

	select {}

}
