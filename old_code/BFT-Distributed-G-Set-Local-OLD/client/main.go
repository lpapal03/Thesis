package main

import (
	"flag"
	"frontend/modules"
	"frontend/tools"
)

func main() {
	tools.ResetLogFile()

	var auto bool
	var clients int
	var reqs int

	flag.BoolVar(&auto, "auto", false, "Automated")
	flag.IntVar(&clients, "clients", 1, "Amount of automated clients")
	flag.IntVar(&reqs, "reqs", 5, "Amount of requests")

	flag.Parse()

	if auto {
		modules.StartAutomated(clients, reqs)
		return
	}
	modules.StartInteractive()
}
