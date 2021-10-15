package main

// 处理操作系统信号
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// 创建将被发送到接收信号的通道
	// 当发送信号且通道未准备好时，Notify 不会阻塞
	// 所以最好创建缓冲通道
	sChan := make(chan os.Signal, 1)

	// Notify 将捕获给定的信号并通过 sChan 发送 os.Signal 值
	// 如果参数中未指定信号，则匹配所有信号
	signal.Notify(sChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// 创建通道去等待信号被处理
	exitChan := make(chan int)
	go func() {
		signal := <-sChan
		switch signal {
		case syscall.SIGHUP:
			fmt.Println("The calling terminal has been closed")
			exitChan <- 0
		case syscall.SIGINT:
			fmt.Println("The process has been interrupted by CTRL+C")
			exitChan <- 1

		case syscall.SIGTERM:
			fmt.Println("kill SIGTERM was executed for process")
			exitChan <- 1

		case syscall.SIGQUIT:
			fmt.Println("kill SIGQUIT was executed for process")
			exitChan <- 1
		}
	}()

	code := <-exitChan
	os.Exit(code)
}