package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Linkedlist struct {
	head *Node
}

func (list *Linkedlist) addNode(value int) {
	newNode := &Node{value: value}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *Linkedlist) printList() {
	if list.head == nil {
		fmt.Println("List is empty!")
		return
	}
	current := list.head
	for current != nil {
		fmt.Printf("%d->", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	root := Linkedlist{}

	root.addNode(10)
	root.addNode(20)
	root.addNode(30)
	root.addNode(40)
	root.addNode(50)
	root.printList()

}
