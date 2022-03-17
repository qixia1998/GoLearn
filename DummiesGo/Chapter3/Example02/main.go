package main

// 使用逻辑于比较运算符
import "fmt"

func main() {

	num := 6
	fmt.Println(num == 0) // 等于 false

	fmt.Println(num != 0) // 不等于 true

	fmt.Println(num < 0) // 小于 false

	fmt.Println(num > 0) // 大于 true

	fmt.Println(num <= 0) // 小于等于 false

	fmt.Println(num >= 0) // 大于等于 true

	// 逻辑与 && 
	condition := num > 2 && num < 9
	fmt.Println(condition) // true

	// 逻辑或 ||
	condition = num > 9 || num < 2
	fmt.Println(condition) // false

	// 逻辑非 !
	condition = !(num > 9 && num < 2)
	fmt.Println(condition) // true
}