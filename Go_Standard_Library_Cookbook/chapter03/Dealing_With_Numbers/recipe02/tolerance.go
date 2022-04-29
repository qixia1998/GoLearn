package main

// 浮点数比较
import (
	"fmt"
	"math"
)

const da = 0.29999999999999998889776975374843459576368331909180
const db = 0.3

func main() {

	daStr := fmt.Sprintf("%.10f", da)
	dbStr := fmt.Sprintf("%.10f", db)

	fmt.Printf("Strings %s = %s equals: %v \n", daStr, dbStr, dbStr == daStr)
	fmt.Printf("Number equals: %v \n", db == da)

	// 由于浮点表示的精度是有限的。对于浮点比较，最好使用具有一定公差的比较。
	fmt.Printf("Number equals with TOLERANCE: %v \n", equals(da, db))

}

const TOLERANCE = 1e-8

// Equals 比较公差为 1e-8 的浮点数
func equals(numA, numB float64) bool {
	delta := math.Abs(numA - numB)
	if delta < TOLERANCE {
		return true
	}
	return false
}
