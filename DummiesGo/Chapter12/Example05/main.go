package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 无缓冲通道
// c := make(chan int, 0)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	fmt.Println("Done and can continue to do other work")
}

func main() {
	s := []int{}
	sliceSize := 10
	for i := 0; i < sliceSize; i++ {
		s = append(s, rand.Intn(100))
	}

	c := make(chan int, 5) // 长度为5的缓冲通道
	partSize := 2
	parts := sliceSize / partSize
	i := 0
	for i < parts {
		go sum(s[i*partSize:(i+1)*partSize], c)
		i += 1
	}

	i = 0
	total := 0
	time.Sleep(1 * time.Second) // 模拟之后的检索
	for i < parts {
		partialSum := <-c // 从通道读取
		fmt.Println("Partial Sum: ", partialSum)
		total += partialSum
		i += 1
	}
	fmt.Println("Total: ", total)
	fmt.Scanln()
}
