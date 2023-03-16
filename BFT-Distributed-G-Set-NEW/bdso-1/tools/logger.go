package tools

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var mu sync.Mutex

func ResetLogFile() {
	// Check if the file exists
	if _, err := os.Stat("log.txt"); err == nil {
		// File exists, delete it
		err := os.Remove("log.txt")
		if err != nil {
			fmt.Println(err)
		}
	}

}

func LogDebug(hostname, event string) {
	log.Println("| " + hostname + " | " + event)
}

func Log(hostname, event string) error {
	LogDebug(hostname, event)
	// Open a log file
	mu.Lock()
	defer mu.Unlock()
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Set the log output to the file
	log.SetOutput(file)

	// Write some log messages
	log.Println("| " + hostname + " | " + event)

	log.SetOutput(os.Stdout)
	return nil
}
