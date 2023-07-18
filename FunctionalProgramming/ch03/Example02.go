package main

import "fmt"

func main() {
	s := "hello"
	if true {
		s := "world"
		fmt.Println(s)
	}
	fmt.Println(s)
}
