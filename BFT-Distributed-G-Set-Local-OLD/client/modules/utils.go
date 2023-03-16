package modules

import (
	"fmt"
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"frontend/tools"
	"strconv"
)

func automated_client_task(id string, req_count int) {
	fmt.Println("ID set to '" + id + "'")
	config.Initialize()
	servers := config.SERVERS
	client := client.CreateClient(id, servers)

	for r := 0; r < req_count; r++ {
		messaging.Add(client, id+"-"+strconv.Itoa(r))
		// waitRandomly(1000, 2000)
		messaging.Get(client)
		// waitRandomly(1000, 2000)
	}
	tools.Log(id, "Done")
}
