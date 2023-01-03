package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// servers := config.SetServers("REMOTE")
	// scenarios.Start(servers, "NORMAL")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf(strings.Split(hostname, ".")[0] + "\n")
	fmt.Println(os.Args[1])

}
