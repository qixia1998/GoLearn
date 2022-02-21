package main

// 按名称打开文件
import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("### File content ###\n%s\n", string(c))
	f.Close()

	f, err = os.OpenFile("temp/test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, "Test string")
	f.Close()
	
}