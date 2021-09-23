package main

import (
	"fmt"
	"time"
)

func recevier(input <- chan int) {
	for i := range input {
		fmt.Println(i)
	}
}

func sender(output chan <- int, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 500)
		output <- i * i
	}
	close(output)
}

func main() {

	ch := make(chan int)

	go sender(ch, 4)
	go recevier(ch)
	time.Sleep(time.Second*5)
	fmt.Println("Done")
}
