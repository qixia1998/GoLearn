package main

import "fmt"

// 使用 if/else 语句

func main() {

	num := 5
	condition := num % 2 == 1
	if condition {
		fmt.Println("Number is odd")
	}
	// 显示检查条件的值
	if condition == true {
		fmt.Println("Number is odd")
	}

	// 使用逻辑表达式作为条件部分
	if num % 2 == 1 {
		fmt.Println("Number is odd")
	}

	// 使用 if/else 语句
	if num % 2 == 1 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}
}