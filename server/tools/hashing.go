package tools

import (
	"crypto/sha512"
	"encoding/hex"
)

func Record_to_string(record string) string {
	h := sha512.New()
	h.Write([]byte(record))
	sha512_hash := hex.EncodeToString(h.Sum(nil))
	return sha512_hash
}
