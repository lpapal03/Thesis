package tools

import (
	"fmt"
	"os"
	"sync"
)

// total add time
// total get time
// total requests
// avg of each
var (
	TOTAL_GET_TIME = 0
	TOTAL_ADD_TIME=0
	REQUESTS=0
)

var counter_mutex sync.Mutex

func saveState() error {
	file, err := os.OpenFile("experiment_results.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "TOTAL_GET_TIME=%d\nTOTAL_ADD_TIME=%d\nREQUESTS=%d\n", BRB_MESSAGES, NORMAL_MESSAGES, BRB_MESSAGES+NORMAL_MESSAGES)
	if err != nil {
		return err
	}

	return nil
}

func IncrementAddTime() {
	counter_mutex.Lock()
	BRB_MESSAGES++
	saveState()
	counter_mutex.Unlock()
}

func IncrementGetTime() {
	counter_mutex.Lock()
	NORMAL_MESSAGES++
	saveState()
	counter_mutex.Unlock()
}
