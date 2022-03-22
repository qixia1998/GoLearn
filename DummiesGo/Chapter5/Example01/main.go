package main

// 定义函数
import (
	"time"
	"fmt"
)

func displayDate() {
	fmt.Println(time.Now().Date())
}

func main() {
	displayDate()
}