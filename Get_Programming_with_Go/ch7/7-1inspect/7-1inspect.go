package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	days := 356.2425
	fmt.Printf("Type %T for %[1]v\n", days)

	a := "text"
	fmt.Printf("Type %T for %[1]v\n", a)

	b := 42
	fmt.Printf("Type %T %[1]v\n", b)

	c := 3.14
	fmt.Printf("Type %T %[1]v\n", c)

	d := true
	fmt.Printf("Type %T %[1]v\n", d)
}
