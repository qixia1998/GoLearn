package main

import (
	"fmt"
	"time"
)

func ShowIt() {
	for {
		time.Sleep(time.Microsecond * 100)
		fmt.Println("Here is it!!!")
	}
}

func main() {

	go ShowIt()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Microsecond * 50)
		fmt.Println("Where is it？")
	}
}
