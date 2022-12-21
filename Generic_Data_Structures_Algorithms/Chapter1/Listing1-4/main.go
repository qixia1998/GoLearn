package main

// Generis

import (
	"fmt"
)

type Stringer = interface {
	String() string
}
type Integer int

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

type String string

func (s String) String() string {
	return string(s)
}

type Student struct {
	Name string
	ID   int
	Age  float64
}

func (s Student) String() string {
	return fmt.Sprintf("%s %d %0.2f", s.Name, s.ID, s.Age)
}
func addStudent[T Stringer](students []T, student T) []T {
	return append(students, student)
}

func main() {
	students := []String{}
	result := addStudent[String](students, "Michael")
	result = addStudent[String](result, "Jennifer")
	result = addStudent[String](result, "Elaine")
	fmt.Println(result)

	students1 := []Integer{}
	result1 := addStudent[Integer](students1, 45)
	result1 = addStudent[Integer](result1, 64)
	result1 = addStudent[Integer](result1, 78)
	fmt.Println(result1)

	students2 := []Student{}
	result2 := addStudent[Student](students2, Student{"John", 213, 17.5})
	result2 = addStudent[Student](result2, Student{"James", 111, 18.75})
	result2 = addStudent(result2, Student{"Marsha", 110, 16.25})
	fmt.Println(result2)

}
