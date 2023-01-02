package modules

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
	"strconv"
	"time"
)

func Leader_task(leader config.Node, peers []config.Node) {
	server := server.Create(leader, peers)
	time.Sleep(time.Second * 1)
	for i := 0; i < 100; i++ {
		m, _ := messaging.StringToMessage([]string{server.Id, messaging.BRACHA_BROADCAST, strconv.Itoa(i)})
		messaging.ReliableBroadcast(server, m)
	}

	for {
		message, _ := server.Receive_socket.RecvMessage(0)
		messaging.HandleMessage(server, message)
	}
}
