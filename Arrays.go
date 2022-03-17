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

// linkedlist

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

	if list.head != nil {

		tempNode.next = list.head
		list.head = &tempNode
		list.len++
	} else {
		list.head = &tempNode
		list.len++
	}
	return
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

func main() {

	t := LinkedList{}

	t.insertBeg(4)
	t.insertBeg(7)

	t.prints()
}
