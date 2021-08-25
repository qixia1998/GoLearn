package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	var length = len(scores)
	for i := 0; i < length; i++ {
		fmt.Printf("% d,", scores[i])
	}
}