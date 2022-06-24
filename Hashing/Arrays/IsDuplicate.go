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

func longestcon(nums []int) int {

	maxNum := 0

	// currentSum := 0

	data := make(map[int]int)

	for windowStart := 0; windowStart <= len(nums)-1; windowStart++ {

		maxNum = data[windowStart]
		data[nums[windowStart]]++

		j := data[nums[windowStart]] - 1

		fmt.Println(data[nums[windowStart]], data, j)
	}
	return data[maxNum]
}

func twoSum(numbers []int, target int) []int {

	start, end := 0, len(numbers)-1
	for start < end {
		currentSum := numbers[start] + numbers[end]
		if currentSum == target {
			return []int{start, end}
		} else if currentSum < target {
			start += 1

		} else {
			end -= 1
			fmt.Print(end)
		}
	}
	return []int{}
}

func main() {

	testArray := make([]int, 0)

	testArray = append(testArray, 2, 7, 11, 15)

	// s := "car"
	// t := "akr"

	// longestcon(testArray)
	twoSum(testArray, 9)
	// fmt.Println(strings.Count("kwam", "a"))
	// h := isAnagramb(s, t)
	// k := IsUniqueBits(s)
	// fmt.Println(k)
	// fmt.Println(h)
	// fmt.Println(IsDuplicateInt(testArray))
}
