package main

import (
	"fmt"
	"sort"
)

func main() {
	// 使用map字面量初始化map
	heights := map[string]int{
		"Peter": 170,
		"Joan":  168,
		"Jan":   175,
	}

	// 检查key 是否存在
	fmt.Println(heights["Jim"]) // 0
	if v, ok := heights["Jim"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist")
	}

	// 删除一个key
	// delete(map, key)
	if _, ok := heights["Joan"]; ok {
		delete(heights, "Joan")
	} else {
		fmt.Println("Key does not exist")
	}

	// 获取map中items的数量
	fmt.Println(len(heights))

	var weights map[string]int
	fmt.Println(len(weights)) // 0

	// map 迭代
	for k, v := range heights {
		fmt.Println(k, v)
	}

	// 获取map所有的keys
	var keys []string
	for k := range heights {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	// 设置map中的迭代顺序

	sort.Strings(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(k, heights[k])
	}
}
