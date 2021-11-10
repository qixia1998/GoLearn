package main

// 将字符串转换为数字
import (
	"fmt"
	"strconv"
)

const bin = "00001"
const hex = "2f"
const intString = "12"
const floatString = "12.3"

func main() {

	// 小数
	res, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed integer: %d\n", res)

	// 解析十六进制
	res64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed hexadecima: %d\n", res64)

	// 解析二进制值
	resBin, err := strconv.ParseInt(bin, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed bin: %d\n", resBin)

	// 解析浮点数
	resFloat, err := strconv.ParseFloat(floatString, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed float: %.5f\n", resFloat)
}