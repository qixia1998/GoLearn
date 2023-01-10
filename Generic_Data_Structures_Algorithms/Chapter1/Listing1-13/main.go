package main

// Race Condition

import (
	"sync"
	"fmt"
)

const (
	number = 1000
)

var countValue int

func main() {
	var wg sync.WaitGroup
	wg.Add(number)

	for i := 0; i < number; i++ {
		go func()  {
			countValue++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("\ncountValue = %d", countValue)
}