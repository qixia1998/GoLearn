package main

// 指定数据类型：显示类型变量
import "fmt"

func main() {
	var num1 = 5 // 类型推断
	var num2 int // 指定类型

	fmt.Println(num1) 
	fmt.Println(num2) // int 类型零值：0

	var num3 float32 // 浮点类型变量
	var name string  // 字符串类型变量
	var raining bool // 布尔类型变量

	fmt.Println(num3) // float32 类型零值：0
	fmt.Println(name) // string 类型零值：""
	fmt.Println(raining) // bool 类型零值：false

	var rates float32 = 4.5 // 声明为 float32 类型然后初始化为 4.5
	fmt.Println(rates) // 4.5
}