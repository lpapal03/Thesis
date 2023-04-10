package modules

import (
	"BFT-Distributed-G-Set-Remote/client"
	"BFT-Distributed-G-Set-Remote/config"
	"BFT-Distributed-G-Set-Remote/messaging"
	"BFT-Distributed-G-Set-Remote/tools"
	"os"
	"strconv"
	"strings"
	"sync"

	zmq "github.com/pebbe/zmq4"
)

func StartAutomated(zctx *zmq.Context, client_count, request_count int) {
	servers := config.Initialize()
	var wg sync.WaitGroup
	wg.Add(client_count)
	for i := 0; i < client_count; i++ {
		host, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		host = strings.Split(host, ".")[0]
		id := host + "_client" + strconv.Itoa(i)
		go func(id string) {
			tools.Log(id, "Id set")
			config.Initialize()
			client := client.CreateClient(id, servers, zctx)
			for r := 0; r < request_count; r++ {
				add_time := messaging.Add(client, id+"_test_"+strconv.Itoa(r))
				_, get_time := messaging.Get(client)
				s := tools.Stats{
					TOTAL_GET_TIME: client.TOTAL_GET_TIME,
					TOTAL_ADD_TIME: client.TOTAL_ADD_TIME,
					REQUESTS:       client.REQUESTS,
				}
				client.TOTAL_ADD_TIME, client.REQUESTS = tools.IncrementAddTime(client.Id, add_time, s)
				client.TOTAL_GET_TIME, client.REQUESTS = tools.IncrementGetTime(client.Id, get_time, s)
			}
			tools.Log(id, "Done")
			wg.Done()
		}(id)
	}
	wg.Wait()
}
