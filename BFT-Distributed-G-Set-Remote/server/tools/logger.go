package tools

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

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

var fileMutex sync.Mutex // declare a mutex for file access

func Log(hostname, event string) error {
	fileMutex.Lock() // acquire the lock

	LogDebug(hostname, event)

	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	now := time.Now().Format("2006/01/02 15:04:05") // format the current time as YYYY/MM/DD HH:MM:SS
	data := fmt.Sprintf("%s | %s | %s", now, hostname, event)

	if _, err = f.WriteString(data); err != nil {
		return err
	}
	f.Close()
	fileMutex.Unlock()
	return nil
}
