package main

// 发现变量类型
import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	// 使用 Printf %T
	fmt.Printf("%T\n", start) // time.Time

	// 使用反射
	fmt.Println(reflect.TypeOf(start)) // time.Time 数据类型
	fmt.Println(reflect.ValueOf(start).Kind()) // struct 数据结构
}