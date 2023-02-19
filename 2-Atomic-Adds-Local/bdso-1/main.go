package main

import (
	"backend/config"
	"backend/modules"
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
	servers := config.SetServerNodes()

	if len(os.Args) < 2 {
		modules.Start(servers, "NORMAL", zctx)
	} else if os.Args[1] == "mutes" || os.Args[1] == "m" {
		modules.Start(servers, "MUTES", zctx)
	}

	select {}
}
