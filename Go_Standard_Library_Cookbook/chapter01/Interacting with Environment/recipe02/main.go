package main

// 访问程序参数
// go build -o test (main.go)
// Linux ./test arg1 arg2 
// Windows test.exe arg1 arg2
import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	// 这个调用将打印所有命令行参数
	fmt.Println(args)

	// 第一个参数, slice的 0 项,是被调用的二进制文件的名称
	programName := args[0]
	fmt.Printf("The binary name is: %s \n", programName)

	// 剩下的参数可以被获取 省略第一个参数.
	otherArgs := args[1:]
	fmt.Println(otherArgs)

	for idx, arg := range otherArgs {
		fmt.Printf("Arg %d = %s \n", idx, arg)
	}
}