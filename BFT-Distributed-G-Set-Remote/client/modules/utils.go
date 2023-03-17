package modules

import (
	"2-Atomic-Adds/config"
	"strings"
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
