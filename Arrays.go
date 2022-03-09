package main

import "fmt"

func TwoS(array []int, target int) []int {

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i] + 1)
	}

	return array

}
func main() {

	var test = []int{7, 8, 3}

	t := TwoS(test, 3)

	// fmt.Println("kkjhghvg", append(t, test...))
	fmt.Println("kkjhghvg", t)
}
