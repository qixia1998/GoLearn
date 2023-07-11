package main

type Person struct {
	Age  int
	Name string
}

func changeName(p *Person, newName string) {
	p.Name = newName
}

// Pure function
func changeNamePure(p Person, newName string) Person {
	return Person{
		Age:  p.Age,
		Name: newName,
	}
}
