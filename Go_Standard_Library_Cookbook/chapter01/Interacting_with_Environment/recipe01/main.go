package main

// 检查Golang的版本
import (
	"log"
	"runtime"
)

const info = `
	Application %s starting.
	The binary was build by Go: %s`

func main() {
	log.Printf(info, "Example", runtime.Version())
}
