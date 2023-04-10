package tools

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	BRB_MESSAGES           = 0
	NORMAL_MESSAGES        = 0
	TOTAL_BRB_TIME         = 0
	COMPLETED_BRB_REQUESTS = 0
)

var counter_mutex sync.Mutex

func saveState(host, port string) error {

	filename := "scenario_results_" + host + "_" + port + ".txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	avg_brb := float64(TOTAL_BRB_TIME) / float64(COMPLETED_BRB_REQUESTS) / float64(time.Millisecond)
	_, err = fmt.Fprintf(file, "BRB_MESSAGES=%d\nNORMAL_MESSAGES=%d\nTOTAL=%d\nAVG_BRB_TIME=%fms\n", BRB_MESSAGES, NORMAL_MESSAGES, BRB_MESSAGES+NORMAL_MESSAGES, avg_brb)
	if err != nil {
		return err
	}

	return nil
}

func IncrementBRBCount(host, port string) {
	counter_mutex.Lock()
	BRB_MESSAGES++
	saveState(host, port)
	counter_mutex.Unlock()
}

func IncrementNormalCount(host, port string) {
	counter_mutex.Lock()
	NORMAL_MESSAGES++
	saveState(host, port)
	counter_mutex.Unlock()
}

func IncrementBRBTime(host, port string, t time.Duration) {
	counter_mutex.Lock()
	TOTAL_BRB_TIME += int(t.Nanoseconds())
	COMPLETED_BRB_REQUESTS++
	saveState(host, port)
	counter_mutex.Unlock()
}
