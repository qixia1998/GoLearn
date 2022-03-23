package main

import (
	"fmt"
	"strconv"
)

// 使用多维数组

func main() {

	// 二维数组
	var table [5][6]string
	for row := 0; row < 5; row++ {
		for column := 0; column < 6; column++ {
			table[row][column] = strconv.Itoa(row) + "," + strconv.Itoa(column)
		}
	}
	fmt.Println(table)

	// 三维数组
	var cube [4][3][3]string
	for row := 0; row < 4; row++ {
		for column := 0; column < 3; column++ {
			for depth := 0; depth < 3; depth++ {
				cube[row][column][depth] = strconv.Itoa(row) + "," + strconv.Itoa(column) + "," + strconv.Itoa(depth)
			}
		}
	}
	fmt.Println(cube)
}