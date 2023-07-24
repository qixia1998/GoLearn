package main

import (
	"fmt"
	"math/rand"
)

func add(a, b int) int {
	return a + b
}

func rollDice() int {
	return rand.Intn(6)
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("dice roll: %v\n", rollDice())
	}
}
