package main

import (
	"bufio"
	"os"
	"time"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	msg := "Rule the world with Golang!!!"
	for _, letter := range msg {
		time.Sleep(time.Millisecond*300)
		writer.WriteByte(byte(letter))
		writer.Flush()
	}
}