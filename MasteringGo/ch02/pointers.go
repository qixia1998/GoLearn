package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)

	// 指向 f 的指针
	fP := &f
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP)
	// f 值的变化
	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)

	// f 的值不会改变
	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)

	// 检查空的结构体
	var k *aStructure

	// 这里是 nil 因为目前 k 指向无处
	fmt.Println(k)

	if k == nil {
		k = new(aStructure)
	}

	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil!")
	}
}
