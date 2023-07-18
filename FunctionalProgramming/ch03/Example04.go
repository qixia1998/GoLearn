package main

import (
	"errors"
	"fmt"
)

func main() {
	//s := "hello"
	//s := "world"
	//fmt.Println(s)

	str1, err := func1()
	if err != nil {
		panic(err)
	}
	str2, err := func2()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v %v\n", str1, str2)
}

func func1() (string, error) {
	return "", errors.New("error1")
}

func func2() (string, error) {
	return "", errors.New("error2")
}
