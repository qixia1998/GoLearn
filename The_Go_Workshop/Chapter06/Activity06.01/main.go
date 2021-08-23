package main

import (
	"fmt"
	"errors"
)

var (
	ErrInvalidLastName = errors.New("invalid last name")
	ErrInvalidRoutingNum = errors.New("invalid routing number") 
)

func main() {
	fmt.Println(ErrInvalidLastName)
	fmt.Println(ErrInvalidRoutingNum)
}