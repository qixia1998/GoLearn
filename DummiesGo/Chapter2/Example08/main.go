package main

import "fmt"

func main() {
	var age int
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%d", &age)
	fmt.Println("Your entered:", age)
}