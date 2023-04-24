package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	fmt.Printf("Please enter the value you want to search: \n")
	var value int
	fmt.Scan(&value)

	var isSearch = false
	var length = len(scores)
	for i := 0; i < length; i++ {
		if scores[i] == value {
			isSearch = true
			fmt.Printf("Found value: %d the index is %d", value, i)
			break
		}
	}
	if !isSearch {
		fmt.Printf("The value was not found: %d", value)
	}
}
