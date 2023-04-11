package main

import (
	"BFT-Distributed-G-Set-Remote/config"
	"BFT-Distributed-G-Set-Remote/modules"
	"BFT-Distributed-G-Set-Remote/tools"
	"flag"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	tools.LOGGING = true
	tools.ResetLogFile()

	wd := "/users/loukis/Thesis/BFT-Distributed-G-Set-Remote"
	_, client_threads := config.GetPortAndThreads(wd + "/config")

	zctx, _ := zmq.NewContext()

	var auto bool
	var reqs int
	var clients int

	flag.BoolVar(&auto, "auto", false, "Automated")
	flag.IntVar(&reqs, "reqs", 5, "Amount of requests")
	flag.IntVar(&clients, "clients", 1, "Amount of clients (if given)")

	flag.Parse()

	if flag.Parsed() {
		if clients != 0 {
			client_threads = clients
		}
	}

	if auto {
		modules.StartAutomated(zctx, client_threads, reqs)
		return
	}
	modules.StartInteractive(zctx)
}
