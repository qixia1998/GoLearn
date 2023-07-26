package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%v\n", add(10, add(10, 5)))
	fmt.Printf("%v\n", add(10, 15))

	fmt.Printf("%v\n", time.Now())
}

func add(a, b int) int {
	return a + b
}
