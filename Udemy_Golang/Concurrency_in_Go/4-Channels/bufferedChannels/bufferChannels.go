package main

import (
	"fmt"
)

func main() {

	c := make(chan string, 3)
	c <- "Hello"
	c <- "Earth"
	c <- "from Mars"
	//c <- "from Venus"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
