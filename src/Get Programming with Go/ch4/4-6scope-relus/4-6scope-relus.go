package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" // era变量在整个包都是可用的

func main() {
	year := 2018                               // era变量和year变量都处于作用域之内
	switch month := rand.Intn(12) + 1; month { // 变量era、year和month都处于作用域之内
	case 2:
		day := rand.Intn(28) + 1 // 变量era、year、month和day都处于作用域之内
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1 // 这两个day变量是全新的变量，跟上面声明的变量的同名变量并不相同
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1 // 这两个day变量是全新的变量，跟上面声明的变量的同名变量并不相同
		fmt.Println(era, year, month, day)
	}
}
