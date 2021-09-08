package main

import (
	"fmt"
	"strings"
)
func main() {
	helloworld := "Hello, World"
	helloMars := strings.Replace(helloworld, "World", "Mars", 1)
	fmt.Println(helloMars)
}
