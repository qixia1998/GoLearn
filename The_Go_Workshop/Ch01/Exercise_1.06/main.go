package main

import (
	"fmt"
	"time"
)

func getCondig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
func main() {
	// Type only
	var start, middle, end float32
	fmt.Println(start, middle, end)
	//Initial value mixed type
	var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	fmt.Println(name, left, right, top, bottom)
	var Debug, LogLevel, startUpTime = getCondig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
//func main() {
//	デバッグ := false
//	日志级别 := "info"
//	ይጀምሩ := time.Now()
//	_A1_Μείγμα := " "
//	fmt.Println(デバッグ, 日志级别, ይጀምሩ, _A1_Μείγμα)
//}
