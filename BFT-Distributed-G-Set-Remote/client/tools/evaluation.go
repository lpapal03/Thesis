package tools

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// total add time
// total get time
// total requests
// avg of each
var (
	TOTAL_GET_TIME = 0
	TOTAL_ADD_TIME = 0
	REQUESTS       = 0
)

var counter_mutex sync.Mutex

func saveState(client_id string) error {

	filename := "scenario_results_" + client_id + ".txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	avg_get := float64(TOTAL_GET_TIME) / float64(REQUESTS) / float64(time.Millisecond)
	avg_add := float64(TOTAL_ADD_TIME) / float64(REQUESTS) / float64(time.Millisecond)
	_, err = fmt.Fprintf(file, "REQUESTS=%d\nAVG_GET_TIME=%fms\nAVG_ADD_TIME=%fms\n", REQUESTS, avg_get, avg_add)

	if err != nil {
		return err
	}

	return nil
}

func IncrementAddTime(client_id string, t time.Duration) {
	counter_mutex.Lock()
	TOTAL_ADD_TIME += int(t.Nanoseconds())
	REQUESTS++
	counter_mutex.Unlock()
	saveState(client_id)
}

func IncrementGetTime(client_id string, t time.Duration) {
	counter_mutex.Lock()
	TOTAL_GET_TIME += int(t.Nanoseconds())
	REQUESTS++
	counter_mutex.Unlock()
	saveState(client_id)
}
