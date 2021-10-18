package main

// 从子进程读取/写入
import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {

	var cmd string

	if runtime.GOOS == "windows" {
		cmd = "dir"
	} else {
		cmd = "ls"
	}

	proc := exec.Command(cmd)

	// 输出将运行过程终止并以字节切片形式返回标准输出。
	buff, err := proc.Output()

	if err != nil {
		panic(err)
	}

	// 子进程的输出以字节切片的形式打印为字符串
	fmt.Println(string(buff))
}
