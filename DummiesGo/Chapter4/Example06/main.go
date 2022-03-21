package main

import "fmt"

// 在for循环中使用标签

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("------------")
	}

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				break
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("-----------")
	}
}
