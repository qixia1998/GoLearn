package main

import "fmt"

type name string

type localtion struct {
	x int
	y int
}

type size struct {
	width int
	height int
}

type dot struct {
	name
	localtion
	size
}

func getDots() []dot {
	var dot1 dot

	dot2 := dot{}
	dot2.name = "A"
	dot2.x = 5
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	dot3 := dot{
		name: "B",
		localtion: localtion{
			x: 13,
			y: 27,
		},
		size: size{
			width: 5,
			height: 7,
		},
	}

	dot4 := dot{}
	dot4.name = "C"
	dot4.localtion.x = 101
	dot4.localtion.y = 209
	dot4.size.width = 87
	dot4.size.height = 43

	return []dot{dot1, dot2, dot3, dot4}
}

func main() {
	dots := getDots()
	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot%v: %#v\n", i+1, dots[i])
	}
}