package main

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data string
	Next *Node
}

func (n *Node) printList() {
	current := n
	for current != nil {
		fmt.Print(current.Data, "->")
		current = current.Next
	}
	fmt.Println("null")
}

func (n *Node) sizeOfList() int {
	c := 0
	current := n
	for current != nil {
		c++
		current = current.Next
	}
	return c
}

func (ll *LinkedList) insertAtBeginning(data string) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
}

func (ll *LinkedList) insertAtEnd(data string) {
	newNode := &Node{Data: data}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func main() {

	nodes := &LinkedList{}
	nodes.insertAtBeginning("Node 1")
	nodes.insertAtBeginning("Node 2")
	nodes.insertAtBeginning("Node 3")
	nodes.Head.printList()
	tam := nodes.Head.sizeOfList()
	fmt.Println("Tam:", tam)

	nodes.insertAtBeginning("Node 4")
	fmt.Println("----------")
	tam2 := nodes.Head.sizeOfList()
	fmt.Println("Tam2:", tam2)
	fmt.Println("----------")
	nodes.Head.printList()
	fmt.Println("----------")
	nodes.insertAtEnd("Node 5")
	nodes.Head.printList()
}
