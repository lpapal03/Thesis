package main

import (
	"backend/config"
	"backend/scenarios"
	"os"
	"runtime/debug"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	debug.SetGCPercent(-1)

	zctx, _ := zmq.NewContext()
	servers := config.SetServerNodes()

	if len(os.Args) < 2 {
		scenarios.Start(servers, "NORMAL", zctx)
	} else if os.Args[1] == "mutes" || os.Args[1] == "m" {
		scenarios.Start(servers, "MUTES", zctx)
	}

	select {}
}
