package modules

import (
	"2-Atomic-Adds/client"
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/messaging"
	"2-Atomic-Adds/tools"
	"strconv"
	"sync"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func StartAutomated(zctx *zmq.Context, client_count, request_count int, network_name string) {
	servers := config.Initialize(network_name)
	var wg sync.WaitGroup
	wg.Add(client_count)
	for i := 0; i < client_count; i++ {

		id := "c" + strconv.Itoa(i)
		go func(id string) {
			tools.Log(id, "Id set")
			config.Initialize(network_name)
			client := client.CreateClient(id, servers, zctx)

			time.Sleep(time.Second * 1)
			for r := 0; r < request_count; r++ {
				messaging.Add(client, id+"-"+strconv.Itoa(r))
				waitRandomly(500, 1000)
				messaging.Get(client)
				waitRandomly(500, 1000)
			}
			tools.Log(id, "Done")
			wg.Done()
		}(id)
	}
	wg.Wait()
}
