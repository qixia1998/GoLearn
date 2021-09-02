package main

import (
	"fmt"
	"math"
)

var Global int = 1234
var AnotherGlobal = - 5678

func main() {
	var j int
	i := Global + AnotherGlobal
	fmt.Println("Initial j value:", j)
	j = Global
	// math.Abs() 需要一个 float64 参数
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d, j=%d, k=%.2f.\n", Global, i, j, k)
}
