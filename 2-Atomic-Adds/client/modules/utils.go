package modules

import (
	"2-Atomic-Adds/config"
	"math/rand"
	"strings"
	"time"
)

func isMessageValid(msg string) bool {
	if msg == "" {
		return false
	}
	if strings.Contains(msg, " ") {
		return false
	}
	if strings.Contains(msg, ".") {
		return false
	}
	if strings.Contains(msg, "{") {
		return false
	}
	if strings.Contains(msg, "}") {
		return false
	}
	if strings.Contains(msg, ";") {
		return false
	}
	return true
}

func isAtomicMessageValid(msg string) bool {
	if msg == "" {
		return false
	}
	if strings.Contains(msg, " ") {
		return false
	}
	if strings.Contains(msg, ".") {
		return false
	}
	if strings.Contains(msg, "{") {
		return false
	}
	if strings.Contains(msg, "}") {
		return false
	}
	parts := strings.Split(msg, ";")
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		if len(p) < 1 {
			return false
		}
	}
	if !config.NetworkExists(strings.Split(msg, ";")[1]) {
		return false
	}
	return true
}

func waitRandomly(min, max int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max - min)
	time.Sleep(time.Duration(min+r) * time.Millisecond)
}
