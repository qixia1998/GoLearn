package main

import (
	"fmt"
	"strconv"
)

// 插值字符串

func main() {
	queue := 5
	name := "Jhon"

	// s := name + ", your queue number is " + queue
	s := name + ", your queue number is " + strconv.Itoa(queue)
	fmt.Println(s)

	s = fmt.Sprintf("%s, your queue number is %d", name, queue)

}