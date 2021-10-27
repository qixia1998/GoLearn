package main

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {

	cmd := "ping"
	timeout := 2 * time.Second

	// 命令行工具 "ping" 被执行 2 秒
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	proc := exec.CommandContext(ctx, cmd, "example.com")

	// 进程输出以 io.ReadCloser 的形式获得。 底层实现使用 os.Pipe
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	// 开始进程
	proc.Start()

	// 为了更舒适地读取，使用 bufio.Scanner。 读取调用被阻塞。
	s := bufio.NewScanner(stdout)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
