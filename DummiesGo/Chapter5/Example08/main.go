package main

import "fmt"

// 使用闭包实现匿名函数

func fib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, (f1 + f2)
		return f1
	}
}

func main() {
	gen := fib()
	//fmt.Println(gen())
	//fmt.Println(gen())

	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
}
