package main

import "fmt"

func main() {
	// index 从0开始
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)

	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}

func sort(arrays []int, length int) {
	for i := 0; i < length; i++ {
		var insertElement = arrays[i] // 选取未排序的新元素
		var insertPosition = i // 插入位置
		for j := insertPosition - 1; j >= 0; j-- {
			// 如果新元素小于已排序的元素，它就向右移动
			if insertElement < arrays[j] {
				arrays[j+1] = arrays[j]
				insertPosition--
			}
		}
		arrays[insertPosition] = insertElement // 插入新元素
	}
}
