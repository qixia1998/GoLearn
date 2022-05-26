package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)
	start := time.Now()
	go doSomething()
	go doSomethingElse()
	wg.Wait()

	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething() {
	time.Sleep(time.Millisecond * 1500)
	fmt.Println("\nI've done something")
	wg.Done()
}

func doSomethingElse() {
	time.Sleep(time.Millisecond * 1500)
	fmt.Println("\nI've done something else")
	wg.Done()
}
