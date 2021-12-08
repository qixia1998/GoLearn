package main

import "fmt"

func main() {
	var scores =[]int{90, 70, 50, 80, 60, 85}

	var length = len(scores)
	var tempArray = make([]int, length+1) // 创建一个新数组

	for i := 0; i <length; i++ {
		tempArray[i] = scores[i]
	}
	tempArray[length] = 75

	scores = tempArray
	for i := 0; i <length+1; i++ {
		fmt.Printf("%d,", scores[i])
	}
}