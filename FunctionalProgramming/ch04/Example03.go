package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 1
	for i < 10 {
		i = add1(i)
		fmt.Printf("%v,", i)
	}
}

func add1(input int) int {
	if input != 0 && input > rand.Intn(input) {
		panic("can not increment any more")
	}
	return input + 1
}
