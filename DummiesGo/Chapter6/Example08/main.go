package main

import (
	"errors"
	"fmt"
)

// 插入 一项 到切片
func insert(orig []int, index int, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New(
			"Index cannot be less than 0")
	}

	if index >= len(orig) {
		return append(orig, value), nil
	}

	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value
	return orig, nil
}

// 从切片删除一项
func delete(orig []int, index int) ([]int, error) {
	if index < 0 || index >= len(orig) {
		return nil, errors.New("Index out of range")
	}
	orig = append(orig[:index], orig[index+1:]...)
	return orig, nil
}

func main() {
	t := []int{1, 2, 3, 4, 5}
	t, err := insert(t, 2, 9)
	if err == nil {
		fmt.Println(t) // 1 2 9 3 4 5]
	} else {
		fmt.Println(err)
	}

	t = []int{1, 2, 3, 4, 5}
	t, err = delete(t, 2)
	if err == nil {
		fmt.Println(t) // [1 2 4 5]
	} else {
		fmt.Println(err)
	}
}
