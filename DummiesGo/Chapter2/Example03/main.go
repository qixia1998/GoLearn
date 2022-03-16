package main

import "fmt"

// 使用短变量声明操作

// num1 := 5 // error: non-declaration statement outside function body
func main() {
	// firstName := "Wei-Meng"

	// firstName, lastName, age := "Wei-Meng", "Lee", 25 // 使用短变量声明初始化多个变量（不同类型）

	// var firstName, lastName string = "Wei-Meng", "Lee"

	// 以下声明是不允许的
	// var firstName, lastName string, age int = "Wei-Meng", "Lee", 25

	// var firstName, lastName, age = "Wei-Meng", "Lee", 25
	var (
		firstName string = "Wei-Meng"
		lastName  string = "Lee"
		age       int    = 25
	)

	fmt.Println(firstName, lastName, age)
}
