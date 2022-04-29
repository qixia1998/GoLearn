package main

import (
	"bytes"
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

	buf := bytes.NewBuffer([]byte{})

	// 实现io.Writer接口的buffer被分配给进程的Stdout
	proc.Stdout = buf

	// 为了在这个例子中避免竞争条件。我们等待直到进程退出。
	proc.Run()

	// 该过程将输出写入缓冲区，我们使用字节来打印输出。
	fmt.Println(string(buf.Bytes()))
}
