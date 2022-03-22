package main

import (
	"fmt"
	"time"
)

// 带有多参数函数定义

func displayDate(format string, prefix string) {
	fmt.Println(prefix, time.Now().Format(format))
}

func main() {
	displayDate("Mon 2006-01-02 15:04:05", "Current Date and Time:")
}