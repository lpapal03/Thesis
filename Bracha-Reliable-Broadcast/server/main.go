package main

import (
	"backend/config"
	"backend/modules"
	"runtime/debug"
)

func main() {
	debug.SetGCPercent(-1)
	servers := config.SetServers("LOCAL")
	modules.Start(servers, "NORMAL")
}
