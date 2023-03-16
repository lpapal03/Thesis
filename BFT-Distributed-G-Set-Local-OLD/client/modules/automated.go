package modules

import (
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"frontend/tools"
	"strconv"
	"sync"
)

func StartAutomated(client_count, request_count int) {
	var wg sync.WaitGroup
	wg.Add(client_count)
	for i := 0; i < client_count; i++ {
		id := "c" + strconv.Itoa(i)
		go func(id string) {
			tools.Log(id, "Id set")
			config.Initialize()
			servers := config.SERVERS
			client := client.CreateClient(id, servers)

			for r := 0; r < request_count; r++ {
				messaging.Add(client, id+"-"+strconv.Itoa(r))
				messaging.Get(client)
			}
			tools.Log(id, "Done")
			wg.Done()
		}(id)
	}
	wg.Wait()
}
