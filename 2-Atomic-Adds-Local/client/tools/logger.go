package tools

import (
	"log"
	"os"
)

func logToFile(r string) {
	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(r)
}

func Log(id, event string) {
	log.Println("| " + id + " | " + event)
	logToFile("| " + id + " | " + event + "\n")
}

func LogFatal(id, event string) {
	log.Fatalln("| " + id + " | " + event)
}
