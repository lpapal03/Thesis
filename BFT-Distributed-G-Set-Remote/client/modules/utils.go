package modules

import (
	"fmt"
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

func truncateResponse(r string, maxCount int) string {
	if maxCount < 0 {
		return r
	}
	records := strings.Split(strings.TrimSpace(r), " ")
	count := len(records)
	if count <= maxCount {
		return r
	}
	truncatedRecords := records[:maxCount]
	omittedCount := count - maxCount
	return strings.Join(truncatedRecords, " ") + fmt.Sprintf(" and %d other records", omittedCount)
}
