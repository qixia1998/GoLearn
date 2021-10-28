package main

import (
	"fmt"
	"sync"
)

func main() {
	// 封装好的计数器
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func ()  {
			defer wg.Done()
			// 执行10万次累加
			for j := 0; j < 100000; j++ {
				counter.Incr() // 收到锁保护的方法
			} 
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

// 线程安全的计数器类型
type Counter struct {
	CounterType int
	Name 		string

	mu 			sync.Mutex
	count 		uint64
}

// 加1的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}