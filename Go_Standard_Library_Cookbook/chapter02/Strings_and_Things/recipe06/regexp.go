package main

import (
	"fmt"
	"regexp"
)

const refString = "Mary had a little lamb"

func main() {
	regexp := regexp.MustCompile("l[a-z]+")
	out := regexp.ReplaceAllString(refString, "replacement")
	fmt.Println(out)
}
