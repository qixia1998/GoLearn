package main

import "fmt"

type Node struct {
	data string
	prev *Node
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "San Francisco"
	head.prev = nil
	head.next = nil

	var nodeOakland *Node = &Node{data: "Oakland", prev: head, next: nil}
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley", prev: nodeOakland, next: nil}
	nodeOakland.next = nodeBerkeley

	tail.data = "Fremont"
	tail.prev = nodeBerkeley
	tail.next = nil
	nodeBerkeley.next = tail
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
	newNode.next = p.next // newNode next point to next node
	p.next = newNode      // current next point to newNode
	newNode.prev = p
	newNode.next.prev = newNode
}

func add(data string) {
	var newNode *Node = new(Node)
	newNode.data = data
	newNode.next = nil
	tail.next = newNode
	newNode.prev = tail
	tail = newNode
}

func output(node *Node) {
	var p = node
	var end *Node = nil
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s->", p.data)
		end = p
		p = p.next
	}
	fmt.Printf("End\n")

	p = end
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s->", p.data)
		p = p.prev
	}
	fmt.Printf("Start\n\n")
}

func main() {
	initial()

	//fmt.Printf("Add a new node Walnut: \n")
	//add("Walnut")
	fmt.Printf("Insert a new node Walnut at index 2 : \n")
	insert(2, "Walnut")

	output(head)
}
