package main

import (
	"DummiesGo/Chapter13/Example02/geometry"
	"fmt"
)

func main() {
	pt1 := geometry.Point{X: 2, Y: 3}
	fmt.Println(pt1)
	fmt.Println(pt1.Length())
}
