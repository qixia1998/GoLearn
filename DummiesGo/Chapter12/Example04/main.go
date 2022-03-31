package main

// 异步等待通道
import (
	"fmt"
	"time"
)

func fib(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
		time.Sleep(2 * time.Second)
	}
	close(c)
}

func counter(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	close(c)
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go fib(10, c1)     // 生成10个斐波那契数列
	go counter(10, c2) // 生成10个数

	//for i := range c1 {
	//	fmt.Println("fib()", i)
	//}
	//
	//for i := range c2 {
	//	fmt.Println("counter()", i)
	//}

	c1Closed := false
	c2Closed := false

	go func() {
		for {
			select {
			case n, ok := <-c1:
				if !ok {
					// 通道关闭并且耗尽
					c1Closed = true
					if c1Closed && c2Closed {
						return
					}
				} else {
					fmt.Println("fib()", n)
				}
			case m, ok := <-c2:
				if !ok {
					// 通道关闭并耗尽
					c2Closed = true
					if c1Closed && c2Closed {
						return
					}
				} else {
					fmt.Println("counter()", m)
				}
			}
		}
	}()
	fmt.Println("Continue to do something else...")
	fmt.Scanln()
}
