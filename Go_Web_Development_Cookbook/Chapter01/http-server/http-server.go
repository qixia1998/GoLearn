package main
// 创建一个简单的HTTP服务
import (
	"fmt"
	"log"
	"net/http"
)
const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}