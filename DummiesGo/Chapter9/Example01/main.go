package main

// JSON 支持的数据类型
// Object String Boolean Number Array null

// 解析JSON 为 struct
import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string
	Lastname  string
}

func main() {
	var person People

	jsonString := `{"firstname": "Wei-Meng", "lastname": "Lee"}`

	err := json.Unmarshal([]byte(jsonString), &person)
	if err == nil {
		fmt.Println(person.Firstname)
		fmt.Println(person.Lastname)
	} else {
		fmt.Println(err)
	}
}
