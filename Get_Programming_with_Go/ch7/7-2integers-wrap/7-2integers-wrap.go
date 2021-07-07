package main

import "fmt"

func main() {
	var red uint8 = 255
	red++
	fmt.Println(red)

	var number int8 = 127
	number++
	fmt.Println(number)

	var red1 uint8 = 255
	red1 += 2
	fmt.Println(red1)

	var number1 int8 = 127
	number1 += 3
	fmt.Println(number1)

	red = 0
	red--
	fmt.Println(red)

	number = -128
	number--
	fmt.Println(number)

	var green uint16 = 65535
	green++
	fmt.Println(green)
}
