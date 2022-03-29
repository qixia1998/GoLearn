package main

// 理解 Goroutine

import (
	"fmt"
	"time"
)

func say(s string, times int) {
	for i := 0; i < times; i++ {
		// 注入100ms 延迟
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

func main() {
	//say("Hello", 3)
	//say("World", 2)
	go say("Hello", 3)
	go say("World", 2)
	fmt.Scanln()
}
