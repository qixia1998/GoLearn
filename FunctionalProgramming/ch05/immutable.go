package main

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

//func main() {
//	p := Person{
//		name: "Benny",
//		age:  55,
//	}
//	setName(&p, "Bjorn")
//	fmt.Println(p.name)
//}
//
//func setName(p *Person, name string) {
//	p.name = name
//}

//func main() {
//	m := map[string]int{}
//	addValue(m, "red", 10)
//	fmt.Printf("%v\n", m)
//}
//
//func addValue(m map[string]int, colour string, value int) {
//	m[colour] = value
//}

//func main() {
//	names := []string{"Miranda"}
//	addValue(names, "Yvonne")
//	fmt.Printf("%v\n", names)
//}
//func addValue(s []string, name string) {
//	s = append(s, name)
//}

//func main() {
//	names := []string{"Miranda"}
//	addValue(&names, "Yvonne")
//	fmt.Printf("%v\n", names)
//}
//func addValue(s *[]string, name string) {
//	*s = append(*s, name)
//}

func immutableCreatePerson() Person {
	p := Person{}
	p = immutableSetName(p, "sean")
	p = immutableSetAge(p, 29)
	return p
}

func immutableSetName(p Person, name string) Person {
	p.name = name
	return p
}

func immutableSetAge(p Person, age int) Person {
	p.age = age
	return p
}

func mutableCreatePerson() *Person {
	p := &Person{}
	mutableSetName(p, "Tom")
	mutableSetAge(p, 31)
	return p
}
func mutableSetName(p *Person, name string) {
	p.name = name
}
func mutableSetAge(p *Person, age int) {
	p.age = age
}
