package main

import "fmt"

// 使用if和初始化语句来保持变量的作用域紧凑

func main() {
	v, err := doSomething()
	if err {
		// 处理错误
	} else {
		fmt.Println(v)
	}

	if v, err := doSomething(); !err {
		fmt.Println(v)
	} else {
		// handle the error
	}
}

func doSomething() (int, bool) {
	//...
	// 返回值的示例
	return 5, false
}
