package main

// 使用for循环
import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 反向输出
	for i := 4; i >= 0; i-- {
		fmt.Println(i)
	}

	// Go不支持预增/预减运算符
	// ++x --x 
}