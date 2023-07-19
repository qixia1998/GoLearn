package main

import "fmt"

func outerFunction() func() {
	fmt.Println("outer function")
	return func() {
		fmt.Println("inner function")
	}
}

func main() {
	//greetingFunc := createGreeting()
	//response := greetingFunc("Ana")
	//fmt.Println(response)

	firstGreeting := createGreeting("Well, hello there")
	secondGreeting := createGreeting("Hola ")
	fmt.Println(firstGreeting("Remi"))
	fmt.Println(firstGreeting("Sean"))
	fmt.Println(secondGreeting("Ana"))
}

//func createGreeting() func(string) string {
//	s := "Hello "
//	return func(name string) string {
//		return s + name
//	}
//}

func createGreeting(greeting string) func(string) string {
	return func(name string) string {
		return greeting + name
	}
}
