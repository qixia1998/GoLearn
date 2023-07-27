package main

import (
	"GoLearn/FunctionalProgramming/ch04/TestableCode/player"
	"math/rand"
)

func main() {
	random := rand.Intn(2)
	player.PlayerSelectPure(random)
}
