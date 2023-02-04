package main

import (
	"fmt"
	"frontend/start"
	"os"
	"strconv"

	zmq "github.com/pebbe/zmq4"
)

func main() {

	zctx, _ := zmq.NewContext()

	if len(os.Args) < 2 {
		start.StartInteractive(zctx)
	}
	if os.Args[1] == "interactive" {
		start.StartInteractive(zctx)
	}
	if (os.Args[1] == "automated" || os.Args[1] == "a") && len(os.Args) < 4 {
		fmt.Println("Not enough arguments")
		return
	}
	if os.Args[1] == "automated" || os.Args[1] == "a" {
		client_count, _ := strconv.Atoi(os.Args[2])
		request_count, _ := strconv.Atoi(os.Args[3])
		start.StartAutomated(zctx, client_count, request_count)
	}
}
