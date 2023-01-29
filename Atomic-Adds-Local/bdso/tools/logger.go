package tools

import "log"

func Log(id, event string) {
	log.Println("| " + id + " | " + event)
}

func LogFatal(id, event string) {
	log.Fatalln("| " + id + " | " + event)
}
