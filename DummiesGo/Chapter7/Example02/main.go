package main

import "fmt"

type point struct {
	x float32
	y float32
	z float32
}

// 返回 指向函数中创建的point结构体的指针
func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}

func main() {

	// pt4 和 pt5 都指向同一个结构体实例
	pt4 := newPoint(7.8, 9.1, 2.3)
	pt5 := pt4
	pt5.x = 0
	fmt.Println(pt4)
	fmt.Println(pt5)

	// pt6 指向point结构体的一个新的实例
	pt6 := *pt4
	pt6.z = 0
	fmt.Println(pt4)
	fmt.Println(pt6)

	// pt7 包含一个pt6副本， pt8指向pt7的副本
	pt7 := pt6
	pt8 := &pt7
	fmt.Println(pt7)
	fmt.Println(pt8)

}
