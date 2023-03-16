package modules

import (
	"BFT-Distributed-G-Set/client"
	"BFT-Distributed-G-Set/messaging"
	"math/rand"
	"strconv"
	"time"
)

func StartAutomated(c *client.Client, req_count int) {
	for i := 0; i < req_count; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(2)
		if n == 0 {
			messaging.Get(c)
		}
		if n == 1 {
			// s := randomString()
			s := c.Hostname + "-test-" + strconv.Itoa(i)
			if isRecordValid(s) {
				messaging.Add(c, s)
			}
		}
	}
}
