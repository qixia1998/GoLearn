package main

import (
	"log"

	"GoLearn/Distributed_Services_with_Go/LetsGo/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
