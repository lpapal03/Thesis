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

	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments")
		return
	}
	if len(os.Args) == 2 {
		start.StartInteractive(zctx, os.Args[1])
	}
	if len(os.Args) == 5 {
		if os.Args[2] == "a" || os.Args[2] == "automated" {
			client_count, _ := strconv.Atoi(os.Args[3])
			request_count, _ := strconv.Atoi(os.Args[4])
			start.StartAutomated(zctx, client_count, request_count, os.Args[1])
		}
	}
	fmt.Println("Invalid arguments")
	return

	// ****************************************START TESTING

	// zctx, _ := zmq.NewContext()
	// config.Initialize("sbdso")
	// servers := config.SERVERS

	// c1 := client.CreateClient("loukas", servers, zctx)
	// c2 := client.CreateClient("marios", servers, zctx)
	// // c3 := client.CreateClient("kostas", servers, zctx)

	// r1 := "marios;bdso-1;hello;world"
	// r2 := "loukas;bdso-2;world;hello"

	// r3 := "marios;bdso-1;1;2"
	// r4 := "loukas;bdso-2;2;1"

	// r5 := "marios;bdso-1;test1;test2"
	// r6 := "loukas;bdso-2;test2;test1"

	// go func() {
	// 	messaging.AddAtomic(c1, r3)
	// 	messaging.AddAtomic(c1, r1)
	// 	messaging.AddAtomic(c1, r5)
	// 	messaging.Get(c1)
	// }()
	// go func() {
	// 	messaging.AddAtomic(c2, r2)
	// 	messaging.AddAtomic(c2, r6)
	// 	messaging.AddAtomic(c2, r4)
	// 	messaging.Get(c2)
	// }()

	// select {}

	// ****************************************END TESTING
}
