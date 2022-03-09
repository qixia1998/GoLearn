package main

// 创建临时文件
import (
	"io/ioutil"
	"os"
	"fmt"
)

func main() {
	tFile, err := ioutil.TempFile("", "gostdcookbook")
	if err != nil {
		panic(err)
	}
	// The called is responsible for handing
	// the clean up.
	defer os.Remove(tFile.Name())

	fmt.Println(tFile.Name())
	
	// TempDir returns
	// The path in string.
	tDir, err := ioutil.TempDir("", "gostdcookbookdir")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tDir)
	fmt.Println(tDir)

}