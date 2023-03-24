package main

import (
	"backend/config"
	"backend/modules"
	"backend/tools"
	"flag"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	tools.ResetLogFile()

	zctx, err := zmq.NewContext()
	if err != nil {
		panic(err)
	}
	server_nodes := config.SetServerNodes()

	var scenario string

	flag.StringVar(&scenario, "s", "NORMAL", "Secnario")

	flag.Parse()

	scenario = strings.ToUpper(scenario)

	modules.Start(server_nodes, scenario, zctx)
}
