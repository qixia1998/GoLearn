package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"Mihalis", "Tsoukalos"}}

	r := reflect.ValueOf(A)
	fmt.Println("String value:", r.String())

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())

		// 检查 Record 中是否嵌套其它结构体
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		// 需要将它转为字符串以方便比较
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}

		// 和之前一样，但是使用了内布值
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}
}
