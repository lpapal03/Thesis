package main

import (
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"frontend/tools"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	tools.ResetLogFile()

	zctx, _ := zmq.NewContext()

	config.Initialize("sbdso")
	servers := config.SERVERS

	c1 := client.CreateClient("loukas", servers, zctx)
	c2 := client.CreateClient("marios", servers, zctx)

	go messaging.AddAtomic(c1, "marios;bdso-1;hello;world")
	go messaging.AddAtomic(c2, "loukas;bdso-2;world;hello")

	select {}

	// if len(os.Args) == 1 {
	// 	fmt.Println("Not enough arguments")
	// 	return
	// }
	// if len(os.Args) == 2 {
	// 	start.StartInteractive(zctx, os.Args[1])
	// 	return
	// }
	// if len(os.Args) == 5 {
	// 	if os.Args[2] == "a" || os.Args[2] == "automated" {
	// 		client_count, _ := strconv.Atoi(os.Args[3])
	// 		request_count, _ := strconv.Atoi(os.Args[4])
	// 		start.StartAutomated(zctx, client_count, request_count, os.Args[1])
	// 	}
	// 	return
	// }
	// fmt.Println("Invalid arguments")
	// return
}
