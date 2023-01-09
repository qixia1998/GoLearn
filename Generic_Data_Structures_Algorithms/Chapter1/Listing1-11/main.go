package main

// quit channel instead of WaitGroup

import (
	"fmt"
	"time"
)

var quit chan bool

func pingGenerator(c chan string) {
	for i := 0; i < 5; i++ {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}

func output(c chan string) {
	for {
		select {
		case value := <- c:
			fmt.Println(value)
		case <-time.After(3 * time.Second):
			fmt.Println("Program timed out.")
			quit <- true
		}
	}
}

func main() {
	quit = make(chan bool)
	c := make(chan string)
	go pingGenerator(c)
	go output(c)
	<- quit
}