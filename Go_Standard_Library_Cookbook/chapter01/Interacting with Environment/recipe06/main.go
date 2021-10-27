package main

// 获取当前进程 PID
import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	pid := os.Getpid()
	fmt.Printf("Process PID: %d \n", pid)

	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}