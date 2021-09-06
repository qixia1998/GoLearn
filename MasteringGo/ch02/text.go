package main
// 文本类型
import "fmt"

func main() {
	aString := "Hello World! €"
	fmt.Println("First character", string(aString[0]))

	// Runes
	// A rune
	r := '€'
	fmt.Println("As an int32 value:", r)
	// Runes 转换到文本
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// 现有字符串作为 runes 打印
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	// 现有字符串打印完字符
	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}
