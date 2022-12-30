package modules

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
	"fmt"
	"time"
)

func Leader_task(leader config.Node, peers []config.Node) {
	server := server.Create(leader, peers)
	// scanner := bufio.NewScanner(os.Stdin)
	// var message string
	// var command string
	time.Sleep(time.Second * 1)
	fmt.Println("Controlling server: " + server.Id)
	m := messaging.StringToMessage([]string{server.Id, messaging.BRACHA_BROADCAST, "123"})
	messaging.ReliableBroadcast(server, m)
	// fmt.Print("Type 'rb' for Reliable Broadcast or 'e' for EXIT\n> ")
	// for scanner.Scan() {
	// 	command = strings.ToLower(scanner.Text())
	// 	if command == "e" {
	// 		return
	// 	}
	// 	if command == "rb" {
	// 		fmt.Print("Message to broadcast\n> ")
	// 		scanner.Scan()
	// 		message = scanner.Text()
	// 		m := messaging.StringToMessage([]string{server.Id, messaging.BRACHA_BROADCAST, message})
	// 		messaging.ReliableBroadcast(server, m)
	// 		return
	// 	}
	// 	if len(command) == 0 {
	// 		fmt.Print("> ")
	// 	} else {
	// 		fmt.Print("Type 'rb' for Reliable Broadcast or 'e' for EXIT\n> ")
	// 	}
	// }

}
