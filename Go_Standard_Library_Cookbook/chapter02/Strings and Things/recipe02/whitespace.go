package main

// 将字符串分解为单词
import (
	"fmt"
	"strings"
)

const refString = "Mary had     a little lamb"

func main() {

	words := strings.Fields(refString)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}