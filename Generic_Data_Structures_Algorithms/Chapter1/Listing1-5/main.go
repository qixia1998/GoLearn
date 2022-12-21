package main

// Unconstrained Generic Type any

import (
	"fmt"
)

type Student struct {
	Name string
	ID   int
	Age  float64
}

func addStudent[T any](students []T, student T) []T {
	return append(students, student)
}

func main() {
	students := []string{}
	result := addStudent[string](students, "Michael")
	result = addStudent[string](result, "Jennifer")
	result = addStudent[string](result, "Elaine")
	fmt.Println(result)

	students1 := []int{}
	result1 := addStudent[int](students1, 45)
	result1 = addStudent[int](result1, 64)
	result1 = addStudent[int](result1, 78)
	fmt.Println(result1)

	students2 := []Student{}
	result2 := addStudent[Student](students2, Student{"John", 213, 17.5})
	result2 = addStudent[Student](result2, Student{"James", 111, 18.75})
	result2 = addStudent(result2, Student{"Marsha", 110, 16.25})
	fmt.Println(result2)
}
