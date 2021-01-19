package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 10  // 声明并初始化
	for count > 0 { // 为循环设置条件
		fmt.Println(count)
		time.Sleep(time.Second)
		count-- // 每次循环之后将计数器的值减一，以免产生无限循环
	}
	fmt.Println("Lifftoff!")
}
