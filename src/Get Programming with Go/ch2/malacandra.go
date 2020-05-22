package main
import "fmt"
func main() {
	var distance = 56000000
	const hoursPerDay = 24

	var days = 28
	fmt.Println(distance / (days *hoursPerDay), "km/h")
}
