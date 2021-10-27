package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func main() {
	cmd := []string{"go", "run", "sample.go"}

	// 命令行工具 "ping" 被执行 2 秒
	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	// 进程输出以 io.ReadCloser 的形式获得。 底层实现使用 os.Pipe
	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	// 出于调试目的，我们观察执行过程的输出
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Program says:", s.Text())
		}
	}()

	// 开始进程
	proc.Start()

	// 现在如下行被写入到子进程标准输入
	fmt.Println("Writing input")
	io.WriteString(stdin, "Hello\n")
	io.WriteString(stdin, "Golang\n")
	io.WriteString(stdin, "is awesome\n")

	time.Sleep(time.Second * 2)

	proc.Process.Kill()

}
