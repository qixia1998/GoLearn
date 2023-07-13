package main

import "fmt"

type phoneNumber string
type Person struct {
	name        string
	phonenumber phoneNumber
}

func (p *Person) setPhoneNumber(s phoneNumber) {
	p.phonenumber = s
}

func (p *Person) update(name string, phonenumber phoneNumber) {
	p.name = name
	p.phonenumber = phonenumber
}

func main() {
	p := Person{
		name:        "John",
		phonenumber: "123",
	}

	fmt.Printf("%v\n", p)
}
