package main

import "fmt"

func main() {
	// 创建一个空数组
	aSlice := []float64{}
	// 因为aSlice为空，长度和容量都是0
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	// 添加元素到一个切片
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice))

	// 一个长度为4的切片
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	// 使用 append
	t = append(t, -5)
	fmt.Println(t)

	// A 2D slice
	// 可以根据需要拥有任意多个维度
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

	// 访问二位切片所有元素
	// 使用双 for 循环
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)
}
