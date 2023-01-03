package main

import (
	"log"

	"GoLearn/proglog/LetsGo/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}