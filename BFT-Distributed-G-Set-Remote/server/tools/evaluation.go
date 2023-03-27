package tools

import "sync"

var (
	BRB_MESSAGES    = 0
	NORMAL_MESSAGES = 0
)

var counter_mutex sync.Mutex

func IncrementBRBCount() {
	counter_mutex.Lock()
	BRB_MESSAGES++
	counter_mutex.Unlock()
}

func IncrementNormalCount() {
	counter_mutex.Lock()
	NORMAL_MESSAGES++
	counter_mutex.Unlock()
}
