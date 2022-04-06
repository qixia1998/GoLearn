package main

// 读取和写入 HTTP headers
import (
	"fmt"
	"net/http"
)

func main() {

	header := http.Header{}

	// 使用header作为切片
	header.Set("Auth-X", "abcdef1234")
	header.Add("Auth-X", "defghijkl")
	fmt.Println(header)

	// 检索header中切片值
	resSlice := header["Auth-X"]
	fmt.Println(resSlice)

	// 得到第一个值
	resFirst := header.Get("Auth-X")
	fmt.Println(resFirst)

	// 用这个替换所有现有值
	header.Set("Auth-X", "newvalue")
	fmt.Println(header)

	// 删除 header
	header.Del("Auth-X")
	fmt.Println(header)
}
