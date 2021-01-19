package main

import "fmt"

func main() {
	var room = "lake"
	switch { // 比较表达式将被放置到单独的分支里面
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough // 下降至下一分支
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}
}
