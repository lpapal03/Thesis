package main

import (
	"backend/config"
	"backend/scenarios"
	"os"
	"runtime/debug"
)

func main() {
	debug.SetGCPercent(-1)
	servers := config.SetServers()

	if len(os.Args) < 2 {
		scenarios.Start(servers, "NORMAL")
	}
	if os.Args[1] == "mutes" || os.Args[1] == "m" {
		scenarios.Start(servers, "MUTES")
	}
}
