package main

// 检索当前工作目录
import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// 可执行文件路径
	fmt.Println(ex)

	// 解析可执行文件目录
	exPath := filepath.Dir(ex)
	fmt.Println("Executable path :" + exPath)

	// 使用 EvalSymlinks 获取实际路径
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlink evaluated:" + realPath)
}