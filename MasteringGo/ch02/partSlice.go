package main

import "fmt"

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	l := len(aSlice)

	// 前 5 个元素
	fmt.Println(aSlice[0:5])

	fmt.Println(aSlice[:5])

	// 最后 2 个元素
	fmt.Println(aSlice[l-2: l])

	fmt.Println(aSlice[l-2:])


	// 前 5 个元素
	t := aSlice[0:5:10]
	fmt.Println(len(t), cap(t))
	// 索引2，3，4 处的元素
	// 容量将是 10 - 2
	t = aSlice[2:5:10]
	fmt.Println(len(t), cap(t))

	// 索引在 0，1，2，3，4 处的元素
	// 新的容量将是 6 - 0
	t = aSlice[:5:6]
	fmt.Println(len(t), cap(t))
}
