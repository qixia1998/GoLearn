package main

// 创建HTTP服务
import (
	"fmt"
	"net/http"
)

type SimpleHTTP struct{}

func (s SimpleHTTP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello world")
}

func main() {
	fmt.Println("Starting HTTP server on port 8080")

	// 最终你可以使用
	//http.ListenAndServe(":8080", SimpleHTTP{})
	s := &http.Server{Addr: ":8080", Handler: SimpleHTTP{}}
	s.ListenAndServe()
}
