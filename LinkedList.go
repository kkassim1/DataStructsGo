package main

import (
	"fmt"
)

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

func main() {

	t := LinkedList{}

	t.insertBeg(4)
	t.insertBeg(7)
	t.insertNth(33, 2)

	t.prints()
}
