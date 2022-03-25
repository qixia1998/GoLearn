package main

// 在结构体中定义方法
import (
	"fmt"
	"math"
)

type point struct {
	x float32
	y float32
	z float32
}

func (p point) length() float64 {
	return math.Sqrt(
		(math.Pow(float64(p.x), 2) +
			math.Pow(float64(p.y), 2) +
			math.Pow(float64(p.z), 2)))
}

func (p *point) move(deltax, deltay, deltaz float32) {
	p.x += deltax
	p.y += deltay
	p.z += deltaz
}

func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}

func main() {
	pt4 := newPoint(7.8, 9.1, 2.3)
	fmt.Println(pt4.length())

	pt4.move(0.1, 0.1, 0.1)
	fmt.Println(*pt4)
}
