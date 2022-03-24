package main

import "fmt"

// 提取数组或者切片的一部分

func main() {
	var c [3]string
	c[0] = "iOS"
	c[1] = "Android"
	c[2] = "Windows"

	fmt.Println(c[0:2]) // [iOS Android]
	fmt.Println(c[:2])  // [iOS Android]
	fmt.Println(c[1:])  // [Android Windows]
	fmt.Println(c[:])   // [iOS Android Windows]

	t := []int{1, 2, 3, 4, 5}
	fmt.Println(len(t)) // 5
	fmt.Println(cap(t)) // 5

	t = t[2:4]
	fmt.Println(t)      // [3 4]
	fmt.Println(len(t)) // 2
	fmt.Println(cap(t)) // 3

	t = t[1:3]
	fmt.Println(t)      // [2 3]
	fmt.Println(len(t)) // 2
	fmt.Println(cap(t)) // 4

	// 遍历切片
	t = []int{1, 2, 3, 4, 5}
	for i, v := range t {
		fmt.Println(i, v)
	}
}
