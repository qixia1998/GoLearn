package main

// 解析JSON 为数组
import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string
	Lastname  string
}

func main() {
	var persons []People
	jsonString :=
		`[
            {
                "firstname":"Wei-Meng",
                "lastname":"Lee"
            },
            {
                "firstname":"Mickey",
                "lastname":"Mouse"
            }
        ]`

	json.Unmarshal([]byte(jsonString), &persons)

	for _, person := range persons {
		fmt.Println(person.Firstname)
		fmt.Println(person.Lastname)
	}
}
