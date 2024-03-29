package main

import "fmt"

func main() {
	// 索引从0开始
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)

	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}

func sort(arrays []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1; j++ {
			if arrays[j] > arrays[j+1] { // 交换
				// var flag = arrays[j]
				// arrays[j] = arrays[j+1]
				// arrays[j+1] = flag
				arrays[j], arrays[j+1] = arrays[j+1], arrays[j]
			}
		}
	}
} 