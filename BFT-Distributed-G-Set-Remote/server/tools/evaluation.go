package tools

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	BRB_MESSAGES    = 0
	NORMAL_MESSAGES = 0
	TOTAL_BRB_TIME  = 0
	BRB_REQUESTS    = 0
)

var counter_mutex sync.Mutex

func saveState() error {
	file, err := os.OpenFile("experiment_results.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	avg_brb := float64(TOTAL_BRB_TIME) / float64(BRB_REQUESTS) / float64(time.Millisecond)
	_, err = fmt.Fprintf(file, "BRB_MESSAGES=%d\nNORMAL_MESSAGES=%d\nTOTAL=%d\nAVG_BRB_TIME=%f", BRB_MESSAGES, NORMAL_MESSAGES, BRB_MESSAGES+NORMAL_MESSAGES, avg_brb)
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

func IncrementBRBTime(t time.Duration) {
	counter_mutex.Lock()
	TOTAL_BRB_TIME += int(t.Nanoseconds())
	BRB_REQUESTS++
	saveState()
	counter_mutex.Unlock()
}
