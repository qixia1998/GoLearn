package main

import "fmt"

func reverse(arrays []int, length int) {
	var middle = length / 2
	for i := 0; i <= middle; i++ {
		//var temp = arrays[i]
		//arrays[i] = arrays[length -i -1]
		//arrays[length -i -1] = temp
		arrays[i], arrays[length -i -1] = arrays[length -i -1], arrays[i]
	}
}

func main() {
	var scores = []int{50, 60, 70, 80, 90}
	var length = len(scores)

	reverse(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}
