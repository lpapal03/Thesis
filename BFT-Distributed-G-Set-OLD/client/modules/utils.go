package modules

import (
	"math/rand"
	"strings"
	"time"
)

func isRecordValid(r string) bool {
	return !(r == "" || strings.TrimSpace(r) == "")
}

func randomString() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(6)
	b := make([]byte, n)
	for i := range b {
		b[i] = 'a' + byte(rand.Intn(26))
	}
	return string(b)
}
