package main

import "fmt"

func main() {
	var count = 0
	for count = 10; count > 10; count-- {
		fmt.Println(count)
	}
	fmt.Println(count) // count变量仍然处于作用域之内
}
