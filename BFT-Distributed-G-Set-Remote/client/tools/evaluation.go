package tools

import (
	"fmt"
	"os"
	"time"
)

type Stats struct {
	TOTAL_GET_TIME int
	TOTAL_ADD_TIME int
	REQUESTS       int
}

func saveState(client_id string, stats Stats) error {

	filename := "scenario_results_" + client_id + ".txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	avg_get := float64(stats.TOTAL_GET_TIME) / float64(stats.REQUESTS) / float64(time.Millisecond)
	avg_add := float64(stats.TOTAL_ADD_TIME) / float64(stats.REQUESTS) / float64(time.Millisecond)
	_, err = fmt.Fprintf(file, "REQUESTS=%d\nAVG_GET_TIME=%fms\nAVG_ADD_TIME=%fms\n", stats.REQUESTS, avg_get, avg_add)

	if err != nil {
		return err
	}

	return nil
}

func IncrementAddTime(client_id string, t time.Duration, stats Stats) (int, int) {
	stats.TOTAL_ADD_TIME += int(t.Nanoseconds())
	stats.REQUESTS++
	saveState(client_id, stats)
	return stats.TOTAL_ADD_TIME, stats.REQUESTS
}

func IncrementGetTime(client_id string, t time.Duration, stats Stats) (int, int) {
	stats.TOTAL_GET_TIME += int(t.Nanoseconds())
	stats.REQUESTS++
	saveState(client_id, stats)
	return stats.TOTAL_GET_TIME, stats.REQUESTS
}
