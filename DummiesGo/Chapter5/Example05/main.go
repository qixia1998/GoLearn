package main

import "fmt"

// 函数返回值
func addNum(num1, num2 int) int {
	return num1 + num2
}

// 函数返回多个值
func countOddEven(s string) (int, int) {
	odds, evens := 0, 0
	for _, c := range s {
		if int(c)%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}

// 命名返回值
func addNumName(num1, num2 int) (sum int) {
	sum = num1 + num2
	return // return sum
}

func main() {
	s := addNum(5, 6)
	fmt.Println(s)

	o, e := countOddEven("12345")
	fmt.Println(o, e)

	_, e = countOddEven("12345")
	fmt.Println(e)
}
