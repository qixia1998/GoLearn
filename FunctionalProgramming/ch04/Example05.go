package main

import "fmt"

var (
	name = "Remi"
)

func sayHello() string {
	return fmt.Sprintf("hello %s", name)
}

//func sayHello(name string) string {
//	return fmt.Sprintf("hello %s", name)
//}

func main() {
	sayHello()
	//sayHello("Remi")
}
