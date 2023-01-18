// This sample program demonstrates how to use an unbuffered
// channel to simulate a move of chess between two goroutines.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quit chan bool

func main() {
	rand.Seed(time.Now().UnixNano())
	move := make(chan int)
	quit = make(chan bool)
	// Launch two players.
	go player("Bobby Fischer", move)
	go player("Boris Spassky", move)
	// Start the move
	move <- 1
	<-quit // Blocks until quit assigned a value
}

// player simulates a person moving in chess.
func player(name string, move chan int) {
	// This function takes data out of the move channel
	// and puts data back into the move channel
	for {
		// Wait for turn to play
		turn := <-move // blocks until move assigned a value (every second)
		// Pick a random number and see if we lose the move
		n := rand.Intn(100)
		if n <= 5 && turn >= 5 {
			fmt.Printf("Player %s was check mated and loses!\n", name)
			quit <- true
			return
		}
		// Display and then increment the total move count by one.
		fmt.Printf("Player %s has moved. Turn %d.\n", name, turn)
		turn++
		// Yield the turn back to the opposing player
		time.Sleep(1 * time.Second)
		move <- turn
	}
}
