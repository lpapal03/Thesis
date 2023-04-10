package tools

import (
	"fmt"
	"os"
	"time"
)

type Stats struct {
	BRB_MESSAGES           int
	NORMAL_MESSAGES        int
	TOTAL_BRB_TIME         int
	COMPLETED_BRB_REQUESTS int
}

func saveState(host, port string, stats Stats) error {

	filename := "scenario_results_" + host + "_" + port + ".txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	avg_brb := float64(stats.TOTAL_BRB_TIME) / float64(stats.COMPLETED_BRB_REQUESTS) / float64(time.Millisecond)
	_, err = fmt.Fprintf(file, "BRB_MESSAGES=%d\nNORMAL_MESSAGES=%d\nTOTAL=%d\nAVG_BRB_TIME=%fms\n", stats.BRB_MESSAGES, stats.NORMAL_MESSAGES, stats.BRB_MESSAGES+stats.NORMAL_MESSAGES, avg_brb)
	if err != nil {
		return err
	}

	return nil
}

func IncrementBRBCount(host, port string, stats Stats) int {
	stats.BRB_MESSAGES++
	saveState(host, port, stats)
	return stats.BRB_MESSAGES
}

func IncrementNormalCount(host, port string, stats Stats) int {
	stats.NORMAL_MESSAGES++
	saveState(host, port, stats)
	return stats.NORMAL_MESSAGES
}

func IncrementBRBTime(host, port string, t time.Duration, stats Stats) (int, int) {
	stats.TOTAL_BRB_TIME += int(t.Nanoseconds())
	stats.COMPLETED_BRB_REQUESTS++
	saveState(host, port, stats)
	return stats.TOTAL_BRB_TIME, stats.COMPLETED_BRB_REQUESTS
}
