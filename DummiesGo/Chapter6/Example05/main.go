package main

import "fmt"

// 创建和初始化切片

func main() {
	t := []int{1, 2, 3, 4, 5}
	fmt.Println(len(t))
	fmt.Println(cap(t))

	// 切片的添加 append()
	t = append(t, 6, 7, 8)
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	t = append(t, 9, 10)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	// u 和 t指向同一个底层数组
	u := t
	fmt.Println(u)
	fmt.Println(t)

	u[9] = 100
	fmt.Println(u)
	fmt.Println(t)

	// t超出容量，指向新的数组
	t = append(t, 11)
	fmt.Println(u)
	fmt.Println(t)

	fmt.Println(len(u))
	fmt.Println(cap(u))
	fmt.Println(len(t))
	fmt.Println(cap(t))
}
