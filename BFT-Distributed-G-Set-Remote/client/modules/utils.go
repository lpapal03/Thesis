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
	records := strings.Split(strings.TrimSpace(r), " ")
	count := len(records)
	if count <= maxCount {
		return r
	}
	truncatedRecords := records[:maxCount]
	omittedCount := count - maxCount
	return strings.Join(truncatedRecords, " ") + fmt.Sprintf(" and %d other records", omittedCount)
}
