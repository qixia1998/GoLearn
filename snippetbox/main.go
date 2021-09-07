package main

import (
	"log"
	"net/http"
)

// 定义一个写入包含字节切片的 home 处理函数
// "Hello from Snippetbox"  作为响应主体
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// 添加一个 showSnippet 处理函数
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// 添加一个 createSnippet 处理函数
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// 使用 http.NewServeMux() 函数去初始化一个新的 servemux,
	// 然后将 home 函数注册为“/” URL 模式的处理程序
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	// 与之前完全相同的方式向 servemux 注册两个新的处理程序函数和相应的 URL 模式。
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// 使用 http.ListenAndServer() 函数去启动一个新的 web 服务。
	// 传入两个参数： 要监听的 TCP 网络地址（在本例中为“:4000”）以及刚刚我们创建的servemux
	// 如果 http.ListenAndServe() 返回错误,我们使用 log.Fatal() 函数记录错误消息并退出。
	// 注意http.ListenAndServe() 返回的任何错误总是非零。
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}