package main

// 在二进制，八进制，十进制和十六进制之间的转换
import (
	"fmt"
	"strconv"
)

const bin = "10111"
const hex = "1A"
const oct = "12"
const dec = "10"
const floatNum = "16.123557"

func main() {

	// 将二进制值转换为十六进制
	v, _ := ConvertInt(bin, 2, 16)
	fmt.Printf("Binary value %s converted to hex: %s\n", bin, v)

	// 将十六进制转换为十进制
	v, _ = ConvertInt(hex, 16, 10)
	fmt.Printf("Hex value %s converted to dec: %s\n", hex, v)

	// 将八进制转换为十六进制
	v, _ = ConvertInt(oct, 8, 16)
	fmt.Printf("Oct value %s converted to hex: %s\n", oct, v)

	// 将十进制转换为八进制
	v, _ = ConvertInt(dec, 10, 8)
	fmt.Printf("Dec value %s converted to oct: %s\n", dec, v)

	//... 类似的任何其它转换也可以完成

}

// ConvertInt converts the given string value of base
// to defined toBase.
func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}