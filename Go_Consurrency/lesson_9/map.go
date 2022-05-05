package main

import (
	"fmt"
	"reflect"
)

func main() {

	m := make(map[string]interface{})
	m["name"] = "qixia"
	m["age"] = 24
	m["addr"] = "ShenZhen"
	m["company"] = "IT"

	for item := range m {
		fmt.Println("item:", reflect.ValueOf(item))
	}
}
