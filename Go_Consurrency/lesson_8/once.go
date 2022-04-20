package main

import (
	"fmt"
	"sync"
)

func Other(o *sync.Once) {
	o.Do(func() {
		fmt.Println("in Other")
	})
}

func main() {
	var once sync.Once

	// 第一个初始化函数
	f1 := func() {
		fmt.Println("in f1")
	}

	once.Do(f1) // 打印 in f1

	// 第二个初始化函数
	f2 := func() {
		fmt.Println("in f2")
	}

	Other(&once) // 无输出
	once.Do(f2)  // 无输出
}
