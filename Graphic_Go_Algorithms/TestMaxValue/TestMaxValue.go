package main

import "fmt"

func max(arrays []int, length int) int {
	for i := 0; i < length-1; i++ {
		if arrays[i] > arrays[i+1] { // 交换
            // var temp = arrays[i]
			// arrays[i] = arrays[i+1]
			// arrays[i+1] = temp
			arrays[i], arrays[i+1] = arrays[i+1], arrays[i]
		}
	}
	var maxValue = arrays[length-1]
	return maxValue
}

func main() {
	var scores = []int{60, 50, 95, 80, 70}
	var length = len(scores)
	var maxValue = max(scores, length)
	fmt.Printf("Max Value = %d\n", maxValue)
}