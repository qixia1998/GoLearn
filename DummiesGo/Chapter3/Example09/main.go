package main

import "fmt"

// 没有控制条件 switch

func main() {

	garde := "C"
	switch garde {
	case "A", "B", "C", "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Undefined")
	}
}
