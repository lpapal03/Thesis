package main

import (
	"backend/config"
	"backend/scenarios"
	"os"

	zmq "github.com/pebbe/zmq4"
)

func main() {

	zctx, err := zmq.NewContext()
	if err != nil {
		panic(err)
	}
	server_nodes, bdso_networks := config.SetServerNodes()
	// s := server.CreateServer(server_nodes[0], server_nodes, zctx, bdso_networks)
	// for i := 0; i < 30; i++ {
	// 	k := strconv.Itoa(i)
	// 	messaging.BdsoAdd(s, "hello"+k, "bdso-1")
	// 	messaging.BdsoAdd(s, "world"+k, "bdso-2")
	// }

	if len(os.Args) < 2 {
		scenarios.Start(server_nodes, "NORMAL", zctx, bdso_networks)
	} else if os.Args[1] == "mutes" || os.Args[1] == "m" {
		scenarios.Start(server_nodes, "MUTES", zctx, bdso_networks)
	}

	select {}
}
