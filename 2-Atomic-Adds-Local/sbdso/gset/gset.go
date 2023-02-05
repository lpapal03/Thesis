package gset

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

// Utility function to convert string record to
// a sha516 string value to be used as key
func string_to_sha512(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	sha512_hash := hex.EncodeToString(h.Sum(nil))
	return sha512_hash
}

// Create gset
func Create() map[string]string {
	return make(map[string]string)
}

// Prints entire gset
func Get(gset map[string]string) {
	for _, value := range gset {
		fmt.Println(value)
	}
}

// Checks if a given record exists in the gset
func Exists(gset map[string]string, record string) bool {
	hash := string_to_sha512(record)
	if _, exists := gset[hash]; exists {
		return true
	}
	return false
}

// Adds record to gset
func Add(gset map[string]string, record string) {
	// create a sha512 value based on the record
	sha512_hash := string_to_sha512(record)
	gset[sha512_hash] = record
}

func GsetToString(gset map[string]string, verbose bool) string {
	if len(gset) == 0 {
		return "{}"
	}
	var s = ""
	if verbose {
		for k, v := range gset {
			s = s + "{key:" + k + ", value:" + v + "},"
		}
	} else {
		for _, v := range gset {
			s = s + "{" + v + "},"
		}
	}
	s = s[:len(s)-1]
	return s
}

// checks for pairs of atomic records. Returns them if they exist.
// atomic message format:
// atomic;sender;signature;destination_network;message
// for it to be atomic:
// 1. Signatures must match
// 2. Senders should be different
// 3. Destination networks should be different
func CheckAtomic(gset map[string]string, signature string) (string, string) {
	for k1, v1 := range gset {
		for k2, v2 := range gset {
			if !strings.Contains(v1, ";") || !strings.Contains(v2, ";") {
				continue
			}
			if k1 == k2 {
				continue
			}
			parts1 := strings.Split(v1, ";")
			parts2 := strings.Split(v2, ";")
			if parts1[0] == "atomic" && parts2[0] == "atomic" && parts1[1] != parts2[1] && parts1[2] == parts2[2] && parts1[2] == signature && parts1[3] != parts2[3] {
				gset[k1] = strings.Replace(v1, "atomic", "atomic-complete", -1)
				gset[k2] = strings.Replace(v2, "atomic", "atomic-complete", -1)
				for _, v := range gset {
					fmt.Println(v)
				}
				return v1, v2
			}

		}
	}
	return "", ""
}
