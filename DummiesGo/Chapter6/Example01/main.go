package main

// 定义一个数组
// 语法： var 数组名 [数组长度] 数据类型

import (
	"fmt"
)

func main() {
	var nums [5]int // 一个整型数组 （5个元素）
	fmt.Println(nums)

	fmt.Println(nums[0]) // 第一个元素
	fmt.Println(nums[1]) // 第二个元素
}