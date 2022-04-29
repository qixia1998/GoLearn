package main

// 从日期中检索时间单位
import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2017, 11, 29, 21, 0, 0, 0, time.Local)
	fmt.Printf("Extracting untils from: %v\n", t)

	dOfMonth := t.Day()
	weekDay := t.Weekday()
	month := t.Month()

	fmt.Printf("The %dth day of %v is %v\n", dOfMonth, month, weekDay)
}
