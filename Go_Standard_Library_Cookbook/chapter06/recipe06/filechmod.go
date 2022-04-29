package main

// 改变文件权限

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Obtain current permissions
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File permission %v\n", fi.Mode())

	// Change permissions
	err = f.Chmod(0777)
	if err != nil {
		panic(err)
	}
	fi, err = f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("File permission %v\n", fi.Mode())
}
