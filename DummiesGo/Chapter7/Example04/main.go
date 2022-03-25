package main

import "fmt"

// 比较结构体

type point1 struct {
	x float32
	y float32
	z float32
}

type point2 struct {
	x    float32
	y    float32
	z    float32
	name []string
}

func main() {
	pt1 := point1{x: 5.6, y: 3.8, z: 6.9}
	pt2 := point1{x: 5.6, y: 3.8, z: 6.9}
	pt3 := point1{x: 6.5, y: 3.8, z: 6.9}

	fmt.Println(pt1 == pt2) // true
	fmt.Println(pt2 == pt3) // false

	//pts1 := point2{x: 5.6, y: 3.8, z: 6.9,
	//	name: []string{"pt1"}}
	//
	//pts2 := point2{x: 5.6, y: 3.8, z: 6.9,
	//	name: []string{"pt2"}}

	// invalid operation: pt1 == pt2 (struct containing
	//   []string cannot be compared)
	//fmt.Println(pts1 == pts2)
}
