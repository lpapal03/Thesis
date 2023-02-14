package main

import (
	"backend/config"
	"backend/scenarios"
	"backend/tools"
	"os"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	tools.ResetLogFile()

	zctx, err := zmq.NewContext()
	if err != nil {
		panic(err)
	}
	server_nodes, bdso_networks := config.SetServerNodes()
	// s1 := server.CreateServer(server_nodes[1], server_nodes, zctx, bdso_networks)
	// s2 := server.CreateServer(server_nodes[2], server_nodes, zctx, bdso_networks)
	// s3 := server.CreateServer(server_nodes[3], server_nodes, zctx, bdso_networks)

	// messaging.BdsoAdd(s1, "hello", "world", "bdso-1", "bdso-2")
	// messaging.BdsoAdd(s2, "hello", "world", "bdso-1", "bdso-2")
	// messaging.BdsoAdd(s3, "hello", "world", "bdso-1", "bdso-2")

	if len(os.Args) < 2 {
		scenarios.Start(server_nodes, "NORMAL", zctx, bdso_networks)
	} else if os.Args[1] == "mutes" || os.Args[1] == "m" {
		scenarios.Start(server_nodes, "MUTES", zctx, bdso_networks)
	}

	select {}
}
