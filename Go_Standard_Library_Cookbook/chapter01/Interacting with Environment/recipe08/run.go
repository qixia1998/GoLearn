package main

// 调用外部进程
import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Run()
	if err != nil {
		fmt.Println(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:\n")
		fmt.Println(out.String())
	}
}
