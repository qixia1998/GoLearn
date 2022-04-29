package main

// 检索子进程信息
import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}
	proc := exec.Command(cmd, "1")
	proc.Start()

	// 没有进程状态被返回，直到进程结束.
	fmt.Printf("Process state for running process: %v\n", proc.ProcessState)

	// PID 可以获取正在运行的进程的事件
	fmt.Printf("PID of running process: %d\n\n", proc.Process.Pid)

}
