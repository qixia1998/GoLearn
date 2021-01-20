package main
import(
	"fmt"
	"time"
)
//func main() {
//	Debug := false
//	LogLevel := "info"
//	startUpTime := time.Now()
//
//	fmt.Println(Debug, LogLevel, startUpTime)
//}
func main() {
	Debug, LogLevel, startUoTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUoTime)
}
