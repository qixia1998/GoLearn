package main

// 创建通道
// ch := make(chan type)
// 使用 <- 将一个值发送到通道
// value := <- ch 接受通道的值
import (
	"fmt"
	"time"
)

// --- 发送数据到一个通道 ---
func sendData(ch chan string) {
	fmt.Println("Sending a string into channel...")
	//time.Sleep(2 * time.Second)
	ch <- "Hello"

	fmt.Println("String has been retrieved from channel...")
}

// --- 从通道中获取数据 ---
func getData(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("String retrieved from channel:", <-ch)
}

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

	fmt.Scanln()
}
