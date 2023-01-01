package main

import (
	"backend/config"
	"backend/modules"
)

func main() {
	servers := config.SetServers("LOCAL")
	modules.Start(servers, "NORMAL")
}
