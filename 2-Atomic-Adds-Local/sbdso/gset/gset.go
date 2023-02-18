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

func atomicExists(gset map[string]string, record string) bool {
	if strings.Contains(record, "atomic") {
		test_record := strings.Replace(record, "atomic", "atomic-complete", 1)
		test_key := string_to_sha512(test_record)
		_, exists := gset[test_key]
		if exists {
			return true
		}
	}
	return false
}

// Checks if a given record exists in the gset
func Exists(gset map[string]string, record string) bool {
	if strings.Contains(record, ".") {
		record = strings.Split(record, ".")[2]
	}
	if atomicExists(gset, record) {
		return true
	}
	hash := string_to_sha512(record)
	if _, exists := gset[hash]; exists {
		return true
	}
	return false
}

// Adds record to gset
func Add(gset map[string]string, record string) {
	if strings.Contains(record, ".") {
		record = strings.Split(record, ".")[2]
	}
	// create a sha512 value based on the record
	if atomicExists(gset, record) {
		return
	}
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
// atomic;sender;peer_id;destination_network;your_message;peer_message
func CheckAtomic(gset map[string]string) (string, string) {
	atomic_found := false
	key1, key2 := "", ""
	for k1, v1 := range gset {
		if atomic_found {
			break
		}
		for k2, v2 := range gset {
			if atomic_found {
				break
			}
			if !strings.Contains(v1, ";") || !strings.Contains(v2, ";") {
				continue
			}
			if k1 == k2 {
				continue
			}
			parts1 := strings.Split(v1, ";")
			parts2 := strings.Split(v2, ";")
			if areAtomic(parts1, parts2) {
				atomic_found = true
				key1, key2 = k1, k2
			}

		}
	}
	if atomic_found {
		v1 := strings.Replace(gset[key1], "atomic", "atomic-complete", 1)
		v2 := strings.Replace(gset[key2], "atomic", "atomic-complete", 1)
		delete(gset, key1)
		delete(gset, key2)
		Add(gset, v1)
		Add(gset, v2)
		return v1, v2
	}
	return "", ""
}

func areAtomic(r1, r2 []string) bool {
	// check tag
	if r1[0] != "atomic" || r2[0] != "atomic" {
		return false
	}
	// check senders
	sender1, peer1 := r1[1], r1[2]
	sender2, peer2 := r2[1], r2[2]
	if sender1 != peer2 || sender2 != peer1 {
		return false
	}
	// check senders
	message1, peer_message1 := r1[4], r1[5]
	message2, peer_message2 := r2[4], r2[5]
	if message1 != peer_message2 || message2 != peer_message1 {
		return false
	}
	return true
}
