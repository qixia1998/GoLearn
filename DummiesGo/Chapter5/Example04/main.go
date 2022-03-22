package main


import "fmt"

// 按值传递参数
func swapbyValue(a, b int) {
	a, b = b, a
}

// 按指针传递参数
func swapbyPointer(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 5
	y := 6
	swapbyValue(x, y)
	fmt.Println(x, y) // 5 6

	swapbyPointer(&x, &y)
	fmt.Println(x, y) // 6 5

}