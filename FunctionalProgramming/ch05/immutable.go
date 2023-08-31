package main

import "fmt"

type Person struct {
	name string
	age  int
}

//func main() {
//	p := Person{
//		name: "Benny",
//		age:  55,
//	}
//	setName(p, "Bjorn")
//	fmt.Println(p.name)
//}
//
//func setName(p Person, name string) {
//	p.name = name
//}

func main() {
	p := Person{
		name: "Benny",
		age:  55,
	}
	setName(&p, "Bjorn")
	fmt.Println(p.name)
}

func setName(p *Person, name string) {
	p.name = name
}
