package main

import (
	"2-Atomic-Adds/tools"
	"fmt"
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
		modules.StartInteractive(zctx, os.Args[1])
	}
	if len(os.Args) == 5 {
		if os.Args[2] == "a" || os.Args[2] == "automated" {
			client_count, _ := strconv.Atoi(os.Args[3])
			request_count, _ := strconv.Atoi(os.Args[4])
			modules.StartAutomated(zctx, client_count, request_count, os.Args[1])
		}
	}
	fmt.Println("Invalid arguments")
	return
}
