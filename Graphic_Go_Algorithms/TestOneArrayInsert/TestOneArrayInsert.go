package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)
	var tempArray = make([]int, length+1)

	insert(scores, length, tempArray, 75, 2)// 插入75到索引为2
	
	scores = tempArray
	
	for i := 0; i < length + 1; i++ {
		fmt.Printf("%d,", scores[i])
	}

}

func insert(array []int, length int, tempArray []int, score int, insertIndex int) {
	for i := 0; i < length; i++ {
		if i < insertIndex {
			tempArray[i] = array[i]
		} else {
			tempArray[i+1] = array[i]
		}
	}
	tempArray[insertIndex] = score
}