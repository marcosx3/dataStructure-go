package array

import "fmt"

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

func main() {

	nodes := &Node{
		Data: "Learning",
		Next: &Node{
			Data: "LinkeD",
			Next: &Node{
				Data: "And Data Struct",
				Next: nil,
			},
		},
	}
	nodes.printList()
}
