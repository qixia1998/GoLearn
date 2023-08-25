package main

import (
	"log"
	"os"
)

func main() {

	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()
}
