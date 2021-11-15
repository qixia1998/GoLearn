package main

// 舍入浮点数
import (
	"fmt"
	"math"
)

var valA float64 = 3.55554444

func main() {

	// 通过将数字转换为整数来舍入数字的错误假设
	intVal := int(valA)
	fmt.Printf("Bad rounding by casting to int: %v\n", intVal)

	fRound := Round(valA)
	fmt.Printf("Rounding by custom function: %v\n", fRound)

}

// Round返回最接近的整数
func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}
