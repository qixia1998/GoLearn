package main

// Adding a New Student by ID Number

import (
	"fmt"
)

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func addStudentID(students []int, student int) []int {
	return append(students, student)
}

func main() {
	students := []string{} // Empty slice
	result := addStudent(students, "Michael")
	result = addStudent(result, "Jennifer")
	result = addStudent(result, "Elaine")
	fmt.Println(result)

	students1 := []int{} // Empty slice
	result1 := addStudentID(students1, 115)
	result1 = addStudentID(result1, 112)
	result1 = addStudentID(result1, 120)
	fmt.Println(result1)
}
