package main

// 带有参数的函数定义
import (
	"fmt"
	"time"
)

func displayDate(format string) {
	fmt.Println(time.Now().Format(format))
}

func main() {
	displayDate("Mon 2006-01-02 15:04:05") 
}