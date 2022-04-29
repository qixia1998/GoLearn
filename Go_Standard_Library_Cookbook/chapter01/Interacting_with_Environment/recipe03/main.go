package main

// 使用flag包创建程序接口
// go build -o util
// ./util -retry 2 -prefix=example -array=1,2
import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// 自定义类型需要去实现 flag.Value 接口能够在 flag.Var 函数中使用

type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (a *ArrayValue) Set(s string) error {
	*a = strings.Split(s, " ")
	return nil
}

func main() {

	// 使用返回指针的方法提取 flag 值
	retry := flag.Int("retry", -1, "Defines max retry count")

	// 使用 XXXVar 函数读取 flag.
	// 在这种情况下变量必须被定义在 flag 之前
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "Logger prefix")

	var arr ArrayValue
	flag.Var(&arr, "array", "Input array to iterate through.")

	// 执行 flag.Parse 函数，去读取已定义变量的flag.
	// 如果没有这个 flag 调用 变量保持为空
	flag.Parse()

	// 与 flags 无关的示例逻辑
	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Retrying connection")
		logger.Printf("Sending array %v\n", arr)
		retryCount++
	}
}
