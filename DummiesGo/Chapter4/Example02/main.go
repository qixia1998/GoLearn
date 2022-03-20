package main

import "fmt"

// 生成斐波那契数列

func main() {
	max := 100
	a, b := 0, 1
	for b <= max {
		fmt.Println(b)
		a, b = b, a+b
	}

	for b <= max {
		fmt.Println(b)
		a, b = b, a+b
	}

	max = 100
	for a, b := 0, 1; b <= max; a, b = b, a+b {
		println(b)
	}
}
