package main

import (
	"fmt"
	"strings"
	"unicode"
)

const email = "Example@domain.com"
const name = "isaac newton"
const upc = "upc"
const i = "i"

const snakeCase = "first_name"

func main() {

	// 为了比较用户输入，有时最好比较相同情况下的输入
	input := "Example@domain.com"
	input = strings.ToLower(input)
	emailToCompare := strings.ToLower(email)
	matches := input == emailToCompare
	fmt.Printf("Email matches: %t\n", matches)

	upcCode := strings.ToUpper(upc)
	fmt.Println("UPPER case: " + upcCode)

	// 这个合成字符有不同的 upper 示例和 title 示例
	str := "ǳ"
	fmt.Printf("%s in upper: %s and title: %s\n",
		str,
		strings.ToUpper(str),
		strings.ToTitle(str))

	// XXXSpecial 函数的使用
	title := strings.ToTitle(i)
	titleTurk := strings.ToTitleSpecial(unicode.TurkishCase, i)
	if title != titleTurk {
		fmt.Printf("ToTitle is different: %#U vs, %#U \n",
			title[0], []rune(titleTurk)[0])
	}

	// 在某些情况下，需要更正输入以防万一
	correctNameCase := strings.Title(name)
	fmt.Println("Corrected name: " + correctNameCase)

	// 使用 Title 和 ToLower 函数将蛇形案例转换为驼色案例。
	firstNameCamel := toCamelCase(snakeCase)
	fmt.Println("Camel case :" + firstNameCamel)

}

func toCamelCase(input string) string {
	titleSpace := strings.Title(strings.Replace(input, "_", " ", -1))
	camel := strings.Replace(titleSpace, " ", "", -1)
	return strings.ToLower(camel[:1] + camel[1:])
}
