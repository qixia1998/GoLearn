package main

import (
	"fmt"
	"sync"
	"time"
)

type SliceQueue struct {
	data []interface{}
	mu	 sync.Mutex
}

func NewQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]interface{}, 0, n)}
}

// Enqueue 把值放在队尾
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Dequeue 移去队头并返回
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v 
}

func main() {
	var wg sync.WaitGroup
	queue := NewQueue(100)
	wg.Add(2)
	go func()  {
		wg.Done()
		for i := 0; i < 100; i++ {
			queue.Enqueue(i)
		}
	}()
	go func()  {
		for {
			v := queue.Dequeue()
			fmt.Println("出队的值是: ", v)
			time.Sleep(50 * time.Microsecond)
		}
	}()

	wg.Wait()
}