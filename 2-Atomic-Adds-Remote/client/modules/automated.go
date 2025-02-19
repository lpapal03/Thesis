package modules

import (
	"2-Atomic-Adds/client"
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/messaging"
	"2-Atomic-Adds/tools"
	"os"
	"strconv"
	"strings"
	"sync"

	zmq "github.com/pebbe/zmq4"
)

func StartAutomated(zctx *zmq.Context, client_count, request_count int, network_name string) {
	servers := config.Initialize(network_name)
	var wg sync.WaitGroup
	wg.Add(client_count)
	for i := 0; i < client_count; i++ {
		host, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		host = strings.Split(host, ".")[0]
		id := host + "_client_" + strconv.Itoa(i)
		go func(id string) {
			tools.Log(id, "Id set")
			config.Initialize(network_name)
			client := client.CreateClient(id, servers, zctx)
			for r := 0; r < request_count; r++ {
				messaging.Add(client, id+"-test-"+strconv.Itoa(r))
				messaging.Get(client)
			}
			tools.Log(id, "Done")
			wg.Done()
		}(id)
	}
	wg.Wait()
}
