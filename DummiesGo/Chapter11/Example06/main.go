package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var balance int64
var mutex = &sync.Mutex{}

func credit(wg *sync.WaitGroup) {
	// 当我们完成时通知 WaitGroup
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		// 以原子的方式添加100到balance
		atomic.AddInt64(&balance, 100)
		fmt.Println("After crediting, balance is", balance)
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func debit(wg *sync.WaitGroup) {
	// 当我们完成时通知 WaitGroup
	defer wg.Done()
	for i := 0; i < 5; i++ {
		mutex.Lock()
		// 以原子的方式从balance扣除-100
		atomic.AddInt64(&balance, -100)
		fmt.Println("After debiting, balance is", balance)
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	balance = 200
	fmt.Println("Initial balance is", balance)

	wg.Add(1)
	go credit(&wg)

	wg.Add(1)
	go debit(&wg)

	wg.Wait()
	fmt.Println("Final balance is", balance)
}
