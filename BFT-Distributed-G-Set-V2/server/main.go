package main

import (
	"backend/config"
	"backend/scenarios"
)

func main() {

	servers := config.SetServers("LOCAL")
	scenarios.Start(servers, "NORMAL")

}
