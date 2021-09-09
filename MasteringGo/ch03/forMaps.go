package main

import "fmt"

func main() {
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "A value"

	// range 也适用于 maps
	for key, v := range aMap {
		fmt.Println("key:", key, "value:", v)
	}

	for _, v := range aMap {
		fmt.Print(" # ", v)
	}
	fmt.Println()
}
