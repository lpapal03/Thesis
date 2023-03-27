package tools

import (
	"fmt"
	"os"
	"sync"
)

var (
	BRB_MESSAGES    = 0
	NORMAL_MESSAGES = 0
)

var counter_mutex sync.Mutex

func saveState() error {
	file, err := os.OpenFile("experiment_results.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "BRB_MESSAGES=%d\nNORMAL_MESSAGES=%d\nTOTAL=%d\n", BRB_MESSAGES, NORMAL_MESSAGES, BRB_MESSAGES+NORMAL_MESSAGES)
	if err != nil {
		return err
	}

	return nil
}

func IncrementBRBCount() {
	counter_mutex.Lock()
	BRB_MESSAGES++
	saveState()
	counter_mutex.Unlock()
}

func IncrementNormalCount() {
	counter_mutex.Lock()
	NORMAL_MESSAGES++
	saveState()
	counter_mutex.Unlock()
}
