package main

// switch fallthrough
import "fmt"

func main() {

	garde := "C"
	switch garde {
	case "A":
		fallthrough
	case "B":
		fallthrough
	case "C":
		fallthrough
	case "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Absent")
	}
	// 匹配多个case
	grade := "C"
	switch grade {
	case "A", "B", "C", "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Undefined")
	}
}
