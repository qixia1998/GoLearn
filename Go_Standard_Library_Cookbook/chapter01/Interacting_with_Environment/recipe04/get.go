package main

// 获取和设置具有默认值的环境变量
// DB_CONN=db:/user@example && go run get.go
import (
	"log"
	"os"
)

func main() {
	connStr := os.Getenv("DB_CONN")
	log.Printf("Connection string: %s\n", connStr)
}
