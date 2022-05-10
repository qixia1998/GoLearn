package main

import (
	"fmt"

	"GoLearn/building-microservices-with-go/chapter1/rpc/client"
	"GoLearn/building-microservices-with-go/chapter1/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
