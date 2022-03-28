package main

import "fmt"

// 定义一个接口
type DigitsCounter interface {
	CountOddEven() (int, int)
}

type DigitString string

// DigitString 实现 DigitsCounter
func (ds DigitString) CountEven() (int, int) {
	odds, evens := 0, 0
	for _, digit := range ds {
		if digit%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}

func main() {
	s := DigitString("123456789")
	fmt.Println(s.CountEven())

	//var d DigitsCounter
	//d = s
	//fmt.Println(d.CountOddEven())
}
