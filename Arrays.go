package main

import (
	"fmt"
	"sort"
)

func TwoNumSum(array []int, target int) []int {

	sort.Ints(array)
	left, right := 0, len(array)-1

	for left < right {
		currSum := array[left] + array[right]
		if currSum == target {
			fmt.Println(currSum)
			return []int{array[left], array[right]}
		} else if currSum < target {
			left += 1

		} else {
			right -= 1

		}
	}

	return array //[]int{}

}

func main() {

	var test = []int{7, 8, 8, 3, 8}

	t := TwoNumSum(test, 14)

	fmt.Println("elements", t)
}
