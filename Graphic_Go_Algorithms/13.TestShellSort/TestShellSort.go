package main

import "fmt"

func swap(array []int, a int, b int) {
	array[a] = array[a] + array[b]
	array[b] = array[a] - array[b]
	array[a] = array[a] - array[b]
}

func shellSort(array []int, length int) {
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			var j = i
			for {
				if j-gap < 0 || array[j] >= array[j-gap] {
					break
				}
				swap(array, j, j-gap)
				j -= gap
			}
		}
	}
}

func main() {
	var scores = []int{9, 6, 5, 8, 0, 7, 4, 3, 1, 2}
	var length = len(scores)

	shellSort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}
