package main

import "fmt"

// 为外层循环使用标签

func main() {
Outerloop1:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				break Outerloop1
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("----------")
	}

Outerloop:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				continue Outerloop
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("----------")
	}
}
