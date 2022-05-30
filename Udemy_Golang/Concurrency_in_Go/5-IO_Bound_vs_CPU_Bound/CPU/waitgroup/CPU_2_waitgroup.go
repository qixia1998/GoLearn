package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(16)
	wg.Add(8)
	start := time.Now()
	go counta()
	go countb()
	go countc()
	go countd()
	go counte()
	go countf()
	go countg()
	go counth()

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func counta() {
	fmt.Println("AAAA is starting ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("AAAA is done")
	wg.Done()

}

func countb() {
	fmt.Println("BBBB is starting ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("BBBB is done")
	wg.Done()

}

func countc() {
	fmt.Println("CCCC is starting ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("CCCC is done")
	wg.Done()

}

func countd() {
	fmt.Println("DDDD is starting ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("DDDD is done")
	wg.Done()

}

func counte() {
	fmt.Println("EEEE is starting ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("EEEE is done")
	wg.Done()

}

func countf() {
	fmt.Println("FFFF is starting ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("FFFF is done")
	wg.Done()

}

func countg() {
	fmt.Println("GGGG is starting ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("GGGG is done")
	wg.Done()

}

func counth() {
	fmt.Println("HHHH is starting ")
	for i := 1; i < 10_000_000_000; i++ {
	}

	fmt.Println("HHHH is done")
	wg.Done()

}
