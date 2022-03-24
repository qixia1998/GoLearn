package main

import "fmt"

// 创建数组或切片的副本

func main() {

	nums1 := [5]int{1, 2, 3, 4, 5}
	nums2 := nums1
	fmt.Println(nums1)
	fmt.Println(nums2)

	nums2[0] = 0
	fmt.Println(nums1)
	fmt.Println(nums2)

	// 创建切片副本 copy() 函数
	t := []int{1, 2, 3, 4, 5}
	v := make([]int, len(t))
	copy(v, t)
	fmt.Println(t)
	fmt.Println(v)

	t = []int{1, 2, 3, 4, 5}
	v = make([]int, 2, 5)
	copy(v, t)
	fmt.Println(v)
	fmt.Println(t)

	v = make([]int, 10)
	copy(v, t)
	fmt.Println(v)
	fmt.Println(t)
}
