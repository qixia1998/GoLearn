package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	fmt.Printf("Please enter the index to be deleted: \n")
	var index int
	fmt.Scan(&index)

	var length = len(scores)
	var tempArray = make([]int, length - 1) // 创建一个新数组
	for i := 0; i < length; i++ {
		if i < index { // 将index前面的数据复制到tempArray前面
			tempArray[i] = scores[i]
		}
		if i > index { //复制index后的数组到tempArray的末尾
			tempArray[i-1] = scores[i]
		}
	}

	scores = tempArray
	for i := 0; i < length - 1; i++ {
		fmt.Printf("%d,", scores[i])
	}
}
