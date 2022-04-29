package main

// 使用writer连接字符串
import (
	"bytes"
	"fmt"
)

func main() {
	srings := []string{"This ", "is ", "even ", "more ", "performant "}
	buffer := bytes.Buffer{}
	for _, val := range srings {
		buffer.WriteString(val)
	}

	fmt.Println(buffer.String())
}
