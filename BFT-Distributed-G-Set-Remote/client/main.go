package main

import (
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/modules"
	"2-Atomic-Adds/tools"
	"flag"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	tools.ResetLogFile()

	wd := "/users/loukis/Thesis/BFT-Distributed-G-Set-Remote"
	_, num_threads := config.GetPortAndThreads(wd + "/config")
	fmt.Println(num_threads)

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
		modules.StartAutomated(zctx, clients, reqs, "servers")
		return
	}
	modules.StartInteractive(zctx, "servers")
}
