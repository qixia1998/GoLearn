package main

import (
	"fmt"
	"time"
)

func main() {

	var value = 97
	var result = 0
	goChan := make(chan int)
	mainChan := make(chan string)
	calculateSquare := func() {
		time.Sleep(time.Second * 3)
		result = value * value
		goChan <- result
	}
	reportResult := func() {
		fmt.Println(value, "squared is", <-goChan)
		mainChan <- "You can quit now. I'm done."
	}

	go calculateSquare()
	go reportResult()
	<-mainChan
}
