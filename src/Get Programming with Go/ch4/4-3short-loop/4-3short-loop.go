package main

import "fmt"

func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	} // 随着循环结束，count变量将不再处于作用域之内
}
