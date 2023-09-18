package main

import (
	"fmt"
	"strconv"
)

func fmap[A, B any](mapFunc func(A) B, sliceA []A) []B {
	sliceB := make([]B, len(sliceA))
	for i, a := range sliceA {
		sliceB[i] = mapFunc(a)
	}
	return sliceB
}

func main() {
	integers := []int{1, 2, 3}
	strings := fmap(strconv.Itoa, integers)
	fmt.Printf("%T transformed to %T - %v\n", integers,
		strings, strings)
}
