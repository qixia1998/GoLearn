package main

import "fmt"

// 定义一个 point 的结构体

type point struct {
	x float32
	y float32
	z float32
}

func main() {

	var pt1 point
	pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 4.2

	fmt.Println(pt1.x)
	fmt.Println(pt1.y)
	fmt.Println(pt1.z)

	// 使用指定字段值创建和初始化结构体
	//pt2 := point{x: 5.6, y: 3.8, z: 6.9}

	// 可省略字段名， 可以按顺序指定所有字段值
	//pt2 = point{5.6 , 3.8, 6.9}

	pt2 := point{
		x: 5.6,
		y: 3.8,
		z: 6.9,
	}
	fmt.Println(pt2)

	// 初始化结构体省略特定字段
	pt3 := point{x: 5.6, y: 3.8}
	fmt.Println(pt3)
}
