package main

// 使用原子计数器修改共享资源
import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var balance int64

func credit() {
	for i := 0; i < 10; i++ {
		// 以原子的方式添加100到balance
		atomic.AddInt64(&balance, 100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func debit() {
	for i := 0; i < 5; i++ {
		// 以原子的方式从balance扣除 -100
		atomic.AddInt64(&balance, -100)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
	balance = 200
	fmt.Println("Initial balance is", balance)
	go credit()
	go debit()
	fmt.Scanln()
	fmt.Println(balance)
}
