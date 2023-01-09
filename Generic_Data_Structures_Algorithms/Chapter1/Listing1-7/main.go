package main

// Using generic map and filter functions
import (
	"fmt"
)

func GenericMap[T1, T2 any](input []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}

func GenericFilter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, val := range input {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	input := []float64{-5.0, -2.0, 4.0, 8.0}
	result1 := GenericMap[float64, float64](input, func(n float64) float64 {
		return n * n
	})
	fmt.Println(result1)

	greaterThanFive := GenericFilter[int]([]int{4, 6, 5, 2, 20, 1, 7},
		func(i int) bool {
			return i > 5
		})
	fmt.Println(greaterThanFive)

	oddNumbers := GenericFilter[int]([]int{4, 6, 5, 2, 20, 1, 7},
		func(i int) bool {
			return i%2 == 1
		})
	fmt.Println(oddNumbers)

	lengthGreaterThan3 := GenericFilter[string]([]string{"hello", "or",
		"the", "maybe"}, func(s string) bool {
		return len(s) > 3
	})
	fmt.Println(lengthGreaterThan3)
}
