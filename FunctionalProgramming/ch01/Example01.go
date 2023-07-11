package main

import "fmt"

type predicate func(int) bool

func main() {
	is := []int{1, 1, 2, 3, 5, 8, 13}
	larger := filter(is, largerThan5)
	fmt.Printf("%v", larger)
}

func filter(is []int, condition predicate) []int {
	out := []int{}
	for _, i := range is {
		if condition(i) {
			out = append(out, i)
		}
	}
	return out
}

func largerThan5(i int) bool {
	return i > 5
}
