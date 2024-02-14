package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "A"
	head.next = nil

	var nodeB *Node = &Node{
		data: "B",
		next: nil,
	}
	head.next = nodeB

	var nodeC *Node = &Node{
		data: "C",
		next: nil,
	}
	nodeB.next = nodeC

	tail.data = "D"
	tail.next = head
	nodeC.next = tail
}

func insert(insertPosition int, data string) {
	var p = head
	var i = 0

	for {
		if p.next == nil || i >= insertPosition-1 {
			break
		}
		p = p.next
		i++
	}
	var newNode *Node = new(Node)
	newNode.data = data
	newNode.next = p.next
	p.next = newNode
}

func output(node *Node) {
	var p = node
	for {
		fmt.Printf("%s->", p.data)
		p = p.next
		if p == head {
			break
		}
	}
	fmt.Printf("%s\n\n", p.data)
}

func main() {
	initial()

	fmt.Printf("Insert a new node E at index = 2:\n")
	insert(2, "E")

	output(head)
}
