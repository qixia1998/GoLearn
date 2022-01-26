package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter Counter
	for i := 0; i < 10; i++ { // 10个reader
		go func()  {
			for {
				res := counter.Count() // 计数器读操作
				fmt.Println("resd的值为:", res)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for { // 一个writer
		counter.Incr() // 计数器写操作
		time.Sleep(time.Second)
	}
}

// 一个线程安全的计数器
type Counter struct {
	mu		sync.RWMutex
	count 	int64
}

// 使用写锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 使用读锁保护
func (c *Counter) Count() int64 {
	c.mu.Lock()
	defer c.mu.RUnlock()
	return c.count
}