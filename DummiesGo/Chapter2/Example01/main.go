package main

// 使用 var 关键字：类型推断变量
import "fmt"

// var num1 = 5  // 函数外部定义变量
func main() {
	var num1 = 5 // 类型推断
	fmt.Println(num1) // 5
}