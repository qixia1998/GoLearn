package main

// 遍历通道
import (
	"fmt"
	"time"
)

func fib(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a // 阻塞直到通道接收到值
		a, b = b, a+b
		time.Sleep(1 * time.Second)
	}
	close(c) // 关闭通道
}

func main() {
	c := make(chan int)
	go fib(10, c)
	for i := range c { // 从通道读取直到通道关闭
		fmt.Println(i)
	}
}
