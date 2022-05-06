package main

import (
	"fmt"
	"strings"
)

// IsUnique determine if a string has all unique chars
// This solution is based on HashTable (Go map) data structure
func IsUnique(s string) bool {
	set := make(map[rune]bool)

	for _, c := range s {
		if _, ok := set[c]; ok {

			return false
		}

		set[c] = true
	}

	return true
}

func IsUniqueBits(s string) bool {

	s = strings.ToLower(s)
	var vector int32
	for _, rune := range s {
		index := rune - 'a'
		mask := int32(1 << index)
		if (vector & mask) == mask {
			return false
		}
		vector = vector | mask
	}
	return true
}

func main() {

	fmt.Println("test")
}
