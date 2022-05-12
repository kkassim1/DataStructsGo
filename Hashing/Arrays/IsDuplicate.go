package main

import (
	"fmt"
	"strings"
)

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/

func IsDuplicateInt(nums []int) bool {

	hash := make(map[int]bool)
	for i := 0; i < len(nums); i++ {

		if hash[nums[i]] {

			fmt.Print(hash[i])
			return true
		}

		hash[nums[i]] = true

	}

	return false

}

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

// func isAnagram(s string, t string) bool {

// }

func isAnagramb(s string, t string) bool {

	lenS := len(s)
	lenT := len(t)

	if lenS != lenT {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {
		anagramMap[string(s[i])]++
	}

	for i := 0; i < lenT; i++ {
		anagramMap[string(t[i])]--
	}

	for i := 0; i < lenS; i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}

	return true
}

func main() {

	// testArray := make([]int, 0)

	// testArray = append(testArray, 1, 2, 3, 4, 4, 99, 5, 6)

	s := "car"
	t := "akr"

	// fmt.Println(strings.Count("kwam", "a"))
	h := isAnagramb(s, t)
	// k := IsUniqueBits(s)
	// fmt.Println(k)
	fmt.Println(h)
	// fmt.Println(IsDuplicateInt(testArray))
}
