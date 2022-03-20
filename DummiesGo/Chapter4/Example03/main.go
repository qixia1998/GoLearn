package main

import (
	"fmt"
	"strings"
)

func main() {

	// 无限循环
	for {
		fmt.Println("Enter QUIT to exit")
		var input string
		fmt.Print("Please enter a string:")
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "QUIT" {
			break
		}
	}

	// 打印1-9中的奇数
	for n := 1; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
