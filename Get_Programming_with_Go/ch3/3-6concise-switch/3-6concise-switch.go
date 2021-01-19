package main

import "fmt"

func main() {
	fmt.Println("There is a cavern entrance here and a path to the east")
	var command = "go inside"

	switch command { // 将命令和给定的多个分支进行比较
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside": // 使用逗号风格多个可选值
		fmt.Println("You find yourself in a dimly lit cavern")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}
}
