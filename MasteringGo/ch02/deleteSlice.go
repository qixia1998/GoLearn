package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1{
		fmt.Println("Need an integer value.")
		return
	}

	index := arguments[1]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Using index", i)

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original slice", aSlice)

	// 删除索引 i 处的元素
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	//  ... 自动扩展 aSlice[i+1:]
	// 元素可以一一附加到 aSlice[:i]
	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("After 1st delection:", aSlice)

	// 删除索引 i 处的元素
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	// 用最后一个元素替换索引为 i 处的元素
	aSlice[i] = aSlice[len(aSlice)-1]
	// 删除最后一个元素
	fmt.Println("After 2nd deletion:", aSlice)
}
