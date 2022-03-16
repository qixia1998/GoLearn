package main

// 类型转换
import (
	"fmt"
	"strconv"
)

func main() {
	// var input string
	// fmt.Print("Please enter your age: ")
	// fmt.Scanf("%s", &input)

	// age, err := strconv.Atoi(input) // 字符串转换为整型
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Your age is: ", age)
	// }

	b, err := strconv.ParseBool("t")
	fmt.Println(b)        // true
	fmt.Println(err)      // nil
	fmt.Printf("%T\n", b) // bool

	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)        // 3.1415
	fmt.Println(err)      // nil
	fmt.Printf("%T\n", f) // float64

	i, err := strconv.ParseInt("-18.56", 10, 64) // base 10, 64-bit
	fmt.Println(i)                               // 0
	fmt.Println(err)                             // strconv.ParseInt: parsing "18.56": invalid syntax
	fmt.Printf("%T\n", i)                        // int64

	u1, err := strconv.ParseUint("18", 10, 64)
	fmt.Println(u1)        // 18
	fmt.Println(err)       // <nil>
	fmt.Printf("%T\n", u1) // uint64

	u2, err := strconv.ParseUint("-18", 10, 64)
	fmt.Println(u2)  // 0
	fmt.Println(err) // strconv.ParseUint: parsing
	// "-18": invalid syntax
	fmt.Printf("%T\n", u2) // uint64

	num1 := 5
	num2 := float32(num1)
	num3 := float64(num2)
	num4 := float32(num3)
	num5 := int(num4)

	fmt.Printf("%T\n", num1) // int
	fmt.Printf("%T\n", num2) // float32
	fmt.Printf("%T\n", num3) // float64
	fmt.Printf("%T\n", num4) // float32
	fmt.Printf("%T\n", num5) // int
}
