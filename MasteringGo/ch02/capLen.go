package main

import "fmt"

func main() {
	// 只定义了长度。容量 = 长度
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))
	// 初始化切片。 容量 = 长度
	b := []int{0, 1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b))

	// 相同的长度和容量
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)

	// 添加一个元素
	aSlice = append(aSlice, 5)

	fmt.Println(aSlice)
	// 容量翻倍
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
	// 现在添加四个元素
	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)

	fmt.Println(aSlice)
	// 容量翻倍
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
}
