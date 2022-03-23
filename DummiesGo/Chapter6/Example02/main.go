package main

import "fmt"

// 数组初始化
func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// 使用 ... 省略数组长度
	nums = [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(nums)) // 获取数组长度
}