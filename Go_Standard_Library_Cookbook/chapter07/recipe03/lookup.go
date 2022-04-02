package main

// 通过IP地址解析域和反之亦然
import (
	"fmt"
	"net"
)

func main() {

	// 通过IP解析
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}

	// 通过地址解析
	ips, err := net.LookupIP("localhost")
	if err != nil {
		panic(err)
	}
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}
