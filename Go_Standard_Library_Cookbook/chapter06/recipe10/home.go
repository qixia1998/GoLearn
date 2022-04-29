package main

// 解析用户主目录
import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The user home directory: " + usr.HomeDir)
}
