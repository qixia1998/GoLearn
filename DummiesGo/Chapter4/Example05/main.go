package main

import "fmt"

// 遍历字符串

func main() {

	// Unicode
	for pos, char := range "Hello, world!" {
		fmt.Println(pos, char)
	}

	for pos, char := range "Hello, world!" {
		fmt.Printf("%d %c\n", pos, char)
	}

	for pos, char := range "こんにちは世界" {
		fmt.Printf("%c at byte location %d\n", char, pos)
	}
}
