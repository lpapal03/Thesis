package main

import (
	"backend/config"
	"backend/scenarios"
	"runtime/debug"
)

func main() {
	debug.SetGCPercent(-1)
	servers := config.SetServers()
	scenarios.Start(servers, "MUTES")
}
