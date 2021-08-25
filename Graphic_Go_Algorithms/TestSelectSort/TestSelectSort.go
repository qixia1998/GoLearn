package main
import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)

	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}

func sort(arrays []int, length int) {
	var minIndex int // 保存所选minIndex的索引

	for i := 0; i < length-1; i++ {
		minIndex = i
		// 保存每个循环的minIndex作为第一个元素
		var minValue = arrays[minIndex]
		for j := i; j < length-1; j++ {
			if minValue > arrays[j+1] {
				minValue = arrays[j+1]
				minIndex = j + 1
			}
		}
		// 如果minIndex改变，则当前最小值与minIndex交换
		if i != minIndex {
			// var temp = arrays[i]
			// arrays[i] = arrays[minIndex]
			// arrays[minIndex] = temp	
			arrays[i], arrays[minIndex] = arrays[minIndex], arrays[i]
		}
	}
	
}