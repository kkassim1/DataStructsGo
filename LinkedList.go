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

func main() {

	t := LinkedList{}

	t.insertBeg(4)
	t.insertBeg(7)

	t.prints()
}
