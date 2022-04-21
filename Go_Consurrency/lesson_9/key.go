package main

import "fmt"

func main() {
	var m = make(map[string]int)
	m["a"] = 0
	fmt.Printf("a=%d; b=%d\n", m["a"], m["b"])

	av, aexisted := m["a"]
	bv, bexisted := m["b"]
	fmt.Printf("a=%d, existed: %t; b=%d, existed: %t\n", av, aexisted, bv, bexisted)
}
