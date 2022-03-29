package main

// 使用 WaitGroup
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var balance int64

func credit(wg *sync.WaitGroup) {
	// 当我们完成时通知 WaitGroup
	defer wg.Done()
	for i := 0; i < 10; i++ {
		// 以原子的方式添加100到balance
		atomic.AddInt64(&balance, 100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func debit(wg *sync.WaitGroup) {
	// 当我们完成时通知 WaitGroup
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// 以原子的方式从balance扣除 -100
		atomic.AddInt64(&balance, -100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	balance = 200
	fmt.Println("Initial balance is", balance)

	wg.Add(1) // 将1添加到 WaitGroup 计数器
	go credit(&wg)

	wg.Add(1) // 将1添加到 WaitGroup 计数器
	go debit(&wg)

	wg.Wait() // 阻塞 直到 WaitGroup 计数器为0
	fmt.Println("Final balance is", balance)
}
