package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	err := os.Link("original.txt", "original_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating symlink")

	err = os.Symlink("original.txt", "original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := os.Lstat("original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link info: %+v", fileInfo)

	err = os.Lchown("original_sym.txt", os.Geteuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}
