package main
// 使用GZIP压缩优化HTTP服务响应
import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)
const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)
func helloworld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloworld)
	err := http.ListenAndServe(CONN_HOST + ":" + CONN_PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}