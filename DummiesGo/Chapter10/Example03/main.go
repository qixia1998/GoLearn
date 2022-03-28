package main

import "fmt"

// 使用空接口
// interface{}

func doSomething(i interface{}) {
	fmt.Println(i)
}

func main() {
	doSomething("Hi!")
	doSomething(3.14)
	doSomething([]int{3, 4})
}
