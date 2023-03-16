package main

import (
	"flag"
	"frontend/modules"
	"frontend/tools"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	tools.ResetLogFile()

	zctx, _ := zmq.NewContext()

	var bdso string
	var auto bool
	var clients int
	var reqs int

	flag.StringVar(&bdso, "net", "", "Bdso network")
	flag.BoolVar(&auto, "auto", false, "Automated")
	flag.IntVar(&clients, "clients", 1, "Amount of automated clients")
	flag.IntVar(&reqs, "reqs", 5, "Amount of requests")

	flag.Parse()

	if auto {
		modules.StartAutomated(zctx, clients, reqs, bdso)
		return
	}
	modules.StartInteractive(zctx, bdso)
}
