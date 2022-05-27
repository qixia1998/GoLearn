package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	//ch <- "hello from main"
	go sendMe(ch)
	//for i := 1; i < 2; i++ {
	fmt.Println(<-ch)
	//}

}

func sendMe(ch chan<- string) {

	time.Sleep(time.Second * 2)
	ch <- "SendMe is done"
}
