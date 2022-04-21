package main

import (
	"fmt"
	"time"
)

// map两种常见错误
type Counter struct {
	Website      string
	Start        time.Time
	PageCounters map[string]string
}

func main() {
	var m map[int]int
	// m := make(map[int]int)
	// m[100] = 100 // 未初始化

	fmt.Println(m[100]) // 0

	var c Counter
	c.Website = "baidu.com"
	//c.PageCounters["/"]++

}
