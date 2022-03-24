package main

import "fmt"

// 创建一个空切片

func main() {
	// 使用make() 函数
	s := make([]int, 5)
	fmt.Println(s)

	// ptr: 指向底层数组地址的指针
	// len: 切片长度
	// cap: 容量，或者数组中允许的最大元素数量

	fmt.Println(len(s)) // 长度
	fmt.Println(cap(s)) // 容量

	// 创建具有特定长度和容量的切片
	r := make([]int, 2, 5)
	fmt.Println(len(r))
	fmt.Println(cap(r))
}
