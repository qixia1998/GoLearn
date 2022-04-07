package main

// 使用第三方包
import (
	"fmt"
	"github.com/hackebrot/turtle"
)

func main() {
	emoji, ok := turtle.Emojis["smiley"]
	if !ok {
		fmt.Println("No emoji found.")
	} else {
		fmt.Println(emoji.Char)
	}
}
