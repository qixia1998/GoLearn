package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node) // the first node called head node

func initial() {
	head.data = "San Francisco"
	head.next = nil

	var nodeOakland *Node = &Node{
		data: "Oakland",
		next: nil,
	}
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{
		data: "Oakland",
		next: nil,
	}
	head.next = nodeBerkeley

	var tail *Node = &Node{
		data: "Fremont",
		next: nil,
	}
	nodeBerkeley.next = tail
}

func output(node *Node) {
	var p = node
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s->", p.data)
		p = p.next
	}
	fmt.Printf("End\n\n")
}

func main() {
	initial()
	output(head)
}
