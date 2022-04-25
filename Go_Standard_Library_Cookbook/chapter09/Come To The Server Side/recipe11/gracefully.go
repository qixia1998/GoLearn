package main

// 优雅地关闭HTTP服务
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})

	srv := &http.Server{Addr: ":8080", Handler: mux}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Server error: %s\n", err)
		}
	}()

	log.Println("Server listening on: " + srv.Addr)

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan // 等待SIGINT
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	<-ctx.Done()
	cancel()
	log.Println("Server gracefully stopped")
}
