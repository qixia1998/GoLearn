package main

import (
	"fmt"
	"sync"
)

func main() {
	// 互斥锁保护器
	var mu sync.Mutex
	// 计数器的值
	var count = 0

	// 辅助变量,用来确定所有的goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}