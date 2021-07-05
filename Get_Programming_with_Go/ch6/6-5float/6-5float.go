package main

import (
	"fmt"
	"math"
)

func main() {
	third := 1.0 / 3.0
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)

	fmt.Println(piggyBank == 0.3)

	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)
}
