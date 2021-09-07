package main

import "fmt"

func change(s []string) {
	s[0] = "Change_function"
}

func main() {
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)

	var S0 = a[0:1]
	fmt.Println(S0)
	S0[0] = "S0"

	var S12 = a[1:3]
	fmt.Println(S12)
	S12[0] = "S12_0"
	S12[1] = "S12_1"

	fmt.Println("a:", a)

	// 对切片的更改 -》 对数组的更改
	change(S12)
	fmt.Println("a:", a)

	// S0 容量
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))

	// 添加四个元素到 S0
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")
	a[0] = "-N1"

	// 改变 S0 容量
	// 不再是同一个底层数组
	S0 = append(S0, "N4")

	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))
	// 这个改变不会影响到 S0
	a[0] = "-N1-"

	// 这个改变不会影响到 S12
	a[1] = "-N2-"

	fmt.Println("S0:", S0)
	fmt.Println("a:", a)
	fmt.Println("S12:", S12)
}

