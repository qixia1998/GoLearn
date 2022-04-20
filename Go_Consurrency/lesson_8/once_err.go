package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

func ErrOne() {
	var once sync.Once
	once.Do(func() {
		once.Do(func() {
			fmt.Println("初始化")
		})
	})
}

func ErrTwo() {
	var once sync.Once
	var googleConn net.Conn // 到Google网站的一个连接

	once.Do(func() {
		// 建立到google.com的连接
		googleConn, _ = net.Dial("tcp", "google.com:80")
	})
	// 发送http请求
	googleConn.Write([]byte("GET / HTTP/1.1\r\nHost: google.com\r\n Accept: */*\r\n\r\n"))
	io.Copy(os.Stdout, googleConn)
}

func main() {
	//ErrOne()
	ErrTwo()
}
