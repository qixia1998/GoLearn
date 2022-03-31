package main

// l连接到远程服务器
import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("HELLO GOPHER!\n"),
	}
}

const addr = "localhost:7070"

func main() {
	s := createServer(addr)
	go s.ListenAndServe()

	// 使用普通TCP连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = io.WriteString(conn, "GET / HTTP/1.1\r\nHOST:localhost:7070\r\n\r\n")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(conn)
	conn.SetDeadline(time.Now().Add(time.Second))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.Shutdown(ctx)
}
