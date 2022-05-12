package main

import (
	"fmt"
)

type Node struct {
	next *Node
	data int
}

type LinkedList struct {
	len  int
	head *Node
}

func (list *LinkedList) insertBeg(data int) {

	tempNode := Node{data: data}

	tempNode.data = data
	tempNode.next = list.head
	list.head = &tempNode
	list.len++

	// return
}

func (list *LinkedList) prints() {
	if list.head == nil {
		return
	}

	temp := list.head

	for temp != nil {
		fmt.Println(temp.data)

		temp = temp.next

	}

}

func (list *LinkedList) insertNth(data int, n int) {

	tempNode := Node{data: data}

	tempNode.data = data
	tempNode.next = nil

	if n == 1 {
		tempNode.next = list.head
		list.head = &tempNode
		return
	}

	tempNode2 := list.head

	for i := 0; i < n-2; i++ {
		tempNode2 = tempNode2.next
	}

	tempNode.next = tempNode2.next
	tempNode2.next = &tempNode

}

// linkedlist ----------------------

func isAnagramb(s string, t string) bool {

	lenS := len(s)
	lenT := len(t)

	if lenS != lenT {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {

		fmt.Println("anagram s", anagramMap[string(s[i])])
		anagramMap[string(s[i])]++
		fmt.Println("anagram s aftr:", anagramMap[string(s[i])])

	}

	for i := 0; i < lenT; i++ {
		fmt.Println("de", anagramMap[string(t[i])])
		fmt.Println("t=", t[i])
		anagramMap[string(t[i])]--
		fmt.Println("anagram s aftr:", anagramMap[string(t[i])])
	}

	for i := 0; i < lenS; i++ {
		if anagramMap[string(s[i])] != anagramMap[string(t[i])] {
			fmt.Println(anagramMap[string(s[i])])
			fmt.Println("s", s[i])
			fmt.Println("t", t[i])

			return false
		}
	}

	return true
}

// brute force
func TwoSum(nums []int, target int) []int {

	slice := make([]int, 0)

	slice = append(slice, nums...)

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if (slice[i] + nums[j]) == target {
				// sum := slice[i] + nums[i]

				fmt.Println(slice[i] + nums[j])

				// fmt.Println("ddddddd")

				// slice[i]

			}
		}

	}
	return slice
}

func main() {

	// t := LinkedList{}

	// t.insertBeg(4)
	// t.insertBeg(7)
	// t.insertNth(33, 2)

	// t.prints()

	slice := make([]int, 0)

	slice = append(slice, 5, 9, 0)

	TwoSum(slice, 10)

}
