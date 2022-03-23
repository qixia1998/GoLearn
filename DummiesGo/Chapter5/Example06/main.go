package main

import "fmt"

// 可变参数函数

func addNums(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 固定参数与可变参数一起使用 可变参数位置必须在最后
func addNumsTwo(total int, nums ...int) int {
	fmt.Printf("%T", nums)
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println("Hello")
	fmt.Println("Hello", "World!")
	fmt.Println("Hello", 123, true)

	fmt.Println(addNums(1, 2, 3, 4, 5))
	fmt.Println(addNums(1, 2, 3))
}
