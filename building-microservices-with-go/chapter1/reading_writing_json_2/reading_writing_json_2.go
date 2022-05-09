package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	// 更改输出字段为 "message"
	Message string `json:"message"`
	// 不输出该字段
	// Author string `json:"-"`
	// 	如果值为空，不会输出字段
	// Date string `json:",omitempty"`
	// 将输出转为字符串并重命名 "id"
	// Id int `json:"id, string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}
