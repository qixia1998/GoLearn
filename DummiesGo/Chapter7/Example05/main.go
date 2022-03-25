package main

// 使用cmp进行结构体比较
import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

type Point struct {
	X    float32
	Y    float32
	Z    float32
	Name []string
}

// 只比较X Y
//func (p1 Point) Equal(p2 Point) bool {
//	if p1.X == p2.X &&
//		p1.Y == p2.Y &&
//		p1.Z == p2.Z {
//		return true
//	}
//	return false
//}

func main() {
	pt1 := Point{X: 5.6, Y: 3.8, Z: 6.9,
		Name: []string{"pt1"}}
	pt2 := Point{X: 5.6, Y: 3.8, Z: 6.9,
		Name: []string{"pt"}}
	pt3 := Point{X: 5.6, Y: 3.8, Z: 6.9,
		Name: []string{"pt"}}

	fmt.Println(cmp.Equal(pt1, pt2)) // false
	fmt.Println(cmp.Equal(pt2, pt3)) // true
}
