package main

// 解析本地IP地址

import (
	"fmt"
	"net"
)

func main() {

	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, interf := range interfaces {
		// 为每个接口解析地址
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}
		fmt.Println(interf.Name)
		for _, add := range addrs {
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("\t%v\n", ip)
			}
		}
	}
}
