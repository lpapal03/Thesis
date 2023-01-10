package main

import (
	"fmt"
	"frontend/start"
	"os"
	"runtime/debug"
	"strconv"
)

func main() {

	debug.SetGCPercent(-1)

	if len(os.Args) < 2 {
		start.StartInteractive()
	}
	if os.Args[1] == "interactive" {
		start.StartInteractive()
	}
	if (os.Args[1] == "automated" || os.Args[1] == "a") && len(os.Args) < 4 {
		fmt.Println("Not enough arguments")
		return
	}
	if os.Args[1] == "automated" || os.Args[1] == "a" {
		client_count, _ := strconv.Atoi(os.Args[2])
		request_count, _ := strconv.Atoi(os.Args[3])
		start.StartAutomated(client_count, request_count)
	}
}
