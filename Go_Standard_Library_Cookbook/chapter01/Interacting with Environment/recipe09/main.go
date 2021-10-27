package main

// 检索子进程信息
import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
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

	// 等待函数将等到进程结束.
	proc.Wait()

	// 进程终止后， *os.ProcessState 包含有关进程运行的简单信息
	fmt.Printf("PID: %d\n", proc.ProcessState.Pid())
	fmt.Printf("Process took: %dms\n", proc.ProcessState.SystemTime()/time.Microsecond)
	fmt.Printf("Exited successfully : %t\n", proc.ProcessState.Success())
}
