package main

import "math"

//func DeclarativeFunction() int {
//	return IntRange(-10, 10).
//		Abs().
//		Filter(func(i int64) bool {
//			return i%2 == 0
//		}).
//		Sum()
//	// result = 60
//}

func iterativeFunction() int {
	sum := 0
	for i := -10; i <= 10; i++ {
		absolute := int(math.Abs(float64(i)))
		if absolute%2 == 0 {
			sum += absolute
		}
	}
	return sum
}
