package main

import (
	"BFT-Distributed-G-Set-Remote/config"
	"BFT-Distributed-G-Set-Remote/modules"
	"BFT-Distributed-G-Set-Remote/tools"
	"flag"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	tools.ResetLogFile()

	wd := "/users/loukis/Thesis/BFT-Distributed-G-Set-Remote"
	_, client_threads := config.GetPortAndThreads(wd + "/config")

	zctx, _ := zmq.NewContext()

	var bdso string
	var auto bool
	var reqs int

	flag.StringVar(&bdso, "net", "", "Bdso network")
	flag.BoolVar(&auto, "auto", false, "Automated")
	flag.IntVar(&reqs, "reqs", 5, "Amount of requests")

	flag.Parse()

	if auto {
		modules.StartAutomated(zctx, client_threads, reqs)
		return
	}
	modules.StartInteractive(zctx)
}
