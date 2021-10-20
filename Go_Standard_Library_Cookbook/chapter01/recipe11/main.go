package main

// 优雅地关闭应用程序
import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var writer *os.File

func main() {

	// 文件作为一个日志文件被打开
	// 我们用这种方式表示资源分配

	var err error
	writer, err = os.OpenFile(fmt.Sprintf("test_%d.log",
			time.Now().Unix()), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	
	// 代码在一个 goroutine 中独立地运行.
	// 如果程序是从外部终止，我们需要通过 closeChan 让 goroutine 知道.
	closeChan := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <- closeChan:
				log.Println("Goroutine closing")
				return
			default:
				log.Println("Writing to log")
				io.WriteString(writer, fmt.Sprintf("Logging access %s\n", time.Now().String()))
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, 
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT)

	// 这是从通知函数发送信号的 sigChan 阻塞读取。
	<-sigChan

	// 接收到信号后，读取通道后面的所有代码都可以视为一种清理。
	// 清理部分
	close(closeChan)
	releaseAllResources()
	fmt.Println("The application shut down gracefully")
}

func releaseAllResources() {
	io.WriteString(writer, "Application releasing all resources\n")
	writer.Close()
}


