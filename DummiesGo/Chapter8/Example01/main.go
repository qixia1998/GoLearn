package main

import "fmt"

// 创建map
// 语法 map[keyType] valueType

var heights map[string]int

func main() {
	// map是引用类型，需要在使用之前先使用make()函数初始化
	heights = make(map[string]int)
	heights["Peter"] = 170
	heights["Joan"] = 168
	heights["Jan"] = 175

	fmt.Println(heights["Peter"])
	fmt.Println(heights["Joan"])
	fmt.Println(heights["Jan"])

}
