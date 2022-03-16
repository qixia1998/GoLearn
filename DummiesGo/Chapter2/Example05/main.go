package main

// 删除未使用的变量
import "fmt"

var num1 = 5

func main() {
	var num1 = 5
	_ = num1
	fmt.Println("Hello, World!")
}