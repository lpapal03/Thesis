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
	servers := config.SetServerNodes()

	if len(os.Args) < 2 {
		scenarios.Start(servers, "NORMAL", zctx)
	} else if os.Args[1] == "mutes" || os.Args[1] == "m" {
		scenarios.Start(servers, "MUTES", zctx)
	}

	select {}
}
