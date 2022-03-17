package main

import "fmt"

// Go 不支持三元运算符

func main() {

	num := 5
	// parity := ""
	// if num%2 == 0 {
	// 	parity = "even"
	// } else {
	// 	parity = "odd"
	// }

	parity := checkParity(num)
	fmt.Println(parity)
}

func checkParity(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}
