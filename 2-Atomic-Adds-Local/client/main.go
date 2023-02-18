package main

import (
	"fmt"
	"frontend/start"
	"frontend/tools"
	"os"
	"strconv"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	tools.ResetLogFile()

	zctx, _ := zmq.NewContext()

	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments")
		return
	}
	if len(os.Args) == 2 {
		start.StartInteractive(zctx, os.Args[1])
		return
	}
	if len(os.Args) == 5 {
		if os.Args[2] == "a" || os.Args[2] == "automated" {
			client_count, _ := strconv.Atoi(os.Args[3])
			request_count, _ := strconv.Atoi(os.Args[4])
			start.StartAutomated(zctx, client_count, request_count, os.Args[1])
		}
		return
	}
	fmt.Println("Invalid arguments")
	return
}
