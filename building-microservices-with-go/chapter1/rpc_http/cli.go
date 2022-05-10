package main

import (
	"fmt"

	"GoLearn/building-microservices-with-go/chapter1/rpc_http/client"
	"GoLearn/building-microservices-with-go/chapter1/rpc_http/server"
)

func main() {
	server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}
