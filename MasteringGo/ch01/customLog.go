package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// 调用 os.OpenFile() 创建用于写入的日志文件，
	// 如果它不存在，或者打开它进行写入
	// 通过在其末尾附加新数据 (os.O_APPEND)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "iLog", log.LstdFlags)
	iLog.Println("Hello there！")
	iLog.Println("Mastering Go 3rd edition!")
}
