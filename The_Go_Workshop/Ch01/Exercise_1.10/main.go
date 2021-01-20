package main

import "fmt"

func main() {
	count := 5
	count += 5
	fmt.Println(count)
	count++
	fmt.Println(count)
	count--
	fmt.Println(count)
	count -= 5
	fmt.Println(count)
	name := "John"
	name += " Smith"
	fmt.Println("Hello,", name)
}
