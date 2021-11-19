package main

// 格式化数字
import (
	"fmt"
)

var integer int64 = 32500
var floatNum float64 = 22000.456

func main() {
	// 打印小数常用的方法
	fmt.Printf("%d \n", integer)

	// 总是显示符号
	fmt.Printf("%+d \n", integer)

	// 打印其它基数 X - 16, o - 8, b - 2, d - 10
	fmt.Printf("%X \n", integer)
	fmt.Printf("%#X \n", integer)

	// 前置用 0 填充
	fmt.Printf("%010d \n", integer)

	// 左侧用空格填充
	fmt.Printf("% 10d \n", integer)

	// 右侧填充
	fmt.Printf("% -10d \n", integer)

	// 打印浮点数
	fmt.Printf("%f \n", floatNum)

	// 精度为5的浮点数
	fmt.Printf("%.5f \n", floatNum)

	// 使用科学计数法的浮点数
	fmt.Printf("%e \n", floatNum)

	// %e 大指数 或否则 %f 浮点数
	fmt.Printf("%g \n", floatNum) 

}