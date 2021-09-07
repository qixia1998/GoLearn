package main

import "fmt"

func main() {
	// Byte slice
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)

	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b)

	// 打印字节切片内容为文本
	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))

	// b 的长度
	fmt.Println("Length of b:", len(b))
}
