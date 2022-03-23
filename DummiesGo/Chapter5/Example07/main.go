package main

import "fmt"

// 匿名函数

func doSomething() int {
	return 5
}

func main() {
	//var i func() int
	//var i func() int
	//i = doSomething
	//fmt.Println(i())

	var i func() int
	i = func() int {
		return 5
	}
	fmt.Println(i())
}
