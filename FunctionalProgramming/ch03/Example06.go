package main

import "fmt"

func threeSum(a, b, c int) int {
	return a + b + c
}

func threeSumCurried(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}

func main() {
	fmt.Println(threeSum(10, 20, 30))
	fmt.Println(threeSumCurried(10)(20)(30))
}
