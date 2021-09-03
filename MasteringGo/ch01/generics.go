package main

import (
	"fmt"
)

func print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v, " ")
	}
	fmt.Println()
}

func main() {
	Ints := []int{1, 2, 3}
	Strings := []string{"One", "Two", "Three"}
	print(Ints)
	print(Strings)
}
