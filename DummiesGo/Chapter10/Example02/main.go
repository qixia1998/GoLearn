package main

import (
	"fmt"
	"math"
)

// 使用 Stringer 接口
type Stringer interface {
	String() string
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Circle struct {
	radius float64
	name   string
}

type Square struct {
	length float64
	name   string
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	Area() float64
}

// Circle 实现 Shape
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// 向不满足接口要求的类型添加方法
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

// Square 实现 Shape
func (s Square) Area() float64 {
	return math.Pow(s.length, 2)
}

func calculateArea(listOfShapes []Shape) {
	for _, shape := range listOfShapes {
		fmt.Println("Area of shape is ", shape.Area())
	}
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v (%d years old)", p.FirstName, p.LastName, p.Age)
}

// 实现多个接口
func (c Circle) String() string {
	return fmt.Sprintf("Area is %v Circumference is %v", c.Area(), c.Circumference())
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// 使用空接口
// interface{}

func doSomething(i interface{}) {
	fmt.Println(i)
}
func main() {
	c1 := Circle{radius: 5, name: "c1"}
	s1 := Square{length: 6, name: "s1"}
	t1 := Triangle{base: 6, height: 8}
	//fmt.Println(c1.Area())
	//fmt.Println(s1.Area())

	fmt.Println(c1)

	shapes := []Shape{c1, s1, t1}
	calculateArea(shapes)

	me := Person{"Wei-Meng", "Lee", 38}
	fmt.Println(me)
}
