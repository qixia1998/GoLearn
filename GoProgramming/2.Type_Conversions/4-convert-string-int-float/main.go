package main

import (
	"strconv"
	"fmt"
)

func main() {
	numberFloat := "2.2"
	valueInt, err := strconv.ParseFloat(numberFloat, 64)
	if err != nil {
		fmt.Print("Error happened.")
	} else {
		if valueInt == 2.2 {
			fmt.Println("Success")
		}
	}
}