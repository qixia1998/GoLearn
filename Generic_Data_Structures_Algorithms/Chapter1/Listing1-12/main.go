package main

// Channel Direction

import (
	"fmt"
	"time"
)

var quit chan bool

func pingGenerator(c chan<- string) {
	// The channel can only be sent to - a generator
	for i := 0; i < 5; i++ {
		c <- "ping"
	}
}

func pongGenerator(c chan<- string) {
	// Information can only be sent to the channel - a generator
	for i := 0; i < 5; i++ {
		c <- "pong"
	}
}

func output(c <-chan string) {
	// Information can only be received from the channel - a consumer
	for {
		time.Sleep(time.Second * 1)
		select {
		case value := <-c:
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
	go pongGenerator(c)
	go output(c)
	<-quit
}
