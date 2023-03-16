package modules

import (
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"frontend/tools"
	"strconv"
	"sync"

	zmq "github.com/pebbe/zmq4"
)

func StartAutomated(zctx *zmq.Context, client_count, request_count int, network_name string) {
	var wg sync.WaitGroup
	wg.Add(client_count)
	for i := 0; i < client_count; i++ {
		id := "c" + strconv.Itoa(i)
		go func(id string) {
			tools.Log(id, "Id set")
			config.Initialize(network_name)
			servers := config.SERVERS
			client := client.CreateClient(id, servers, zctx)
			for r := 0; r < request_count; r++ {
				messaging.Add(client, id+"-test-"+strconv.Itoa(r))
				messaging.Get(client)
				// tools.Log(client.Id, r)
			}
			tools.Log(id, "Done")
			wg.Done()
		}(id)
	}
	wg.Wait()
}
