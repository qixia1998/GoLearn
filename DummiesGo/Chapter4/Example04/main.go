package main

import "fmt"

// 遍历数组/切片

func main() {
	var OS [3]string
	OS[0] = "iOS"
	OS[1] = "Android"
	OS[2] = "Windows"

	// i 索引， v 值
	for i, v := range OS {
		fmt.Println(i, v)
	}

	for _, v := range OS {
		fmt.Println(v)
	}

	for i, _ := range OS {
		fmt.Println(i)
	}

	for i := range OS {
		fmt.Println(i)
	}
}
