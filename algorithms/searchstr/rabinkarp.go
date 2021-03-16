package searchnum

import (
	"hash/fnv"
)

// https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm

func RabinKarp(data string, s string) int {
	// this version of rabin-karp algorithm is slower than naive
	// string search due to hashing algorithm used is not rolling hash
	// for proper implementation rabin fingerprint hash or similar should be used
	var fnvHash = fnv.New32()

	if n, err := fnvHash.Write([]byte(s)); n == 0 || err != nil {
		return -1
	}

	pat := fnvHash.Sum32()

	for j := 0; j+len(s) <= len(data); j++ {
		fnvHash.Reset()
		if n, err := fnvHash.Write([]byte(data[j : j+len(s)])); n == 0 || err != nil {
			continue
		}
		if fnvHash.Sum32() == pat {
			if data[j:j+len(s)] == s {
				return j
			}
		}
	}

	return -1
}
