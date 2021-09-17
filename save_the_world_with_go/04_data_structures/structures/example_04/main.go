package main

import "fmt"

type Coordinates struct {
	x int
	t int
}

type Circle struct {
	center Coordinates
	radius int
}

func main() {
	c := Circle{Coordinates{1, 2}, 3}
	fmt.Printf("%+v\n", c)
}
