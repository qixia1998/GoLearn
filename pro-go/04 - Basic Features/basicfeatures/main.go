package main

import (
	"fmt"
	"sort"
)

func main() {
	names := [3]string{"Alice", "Charlie", "Bob"}

	secondName := names[1]

	fmt.Println(secondName)

	sort.Strings(names[:])

	fmt.Println(secondName)
}
