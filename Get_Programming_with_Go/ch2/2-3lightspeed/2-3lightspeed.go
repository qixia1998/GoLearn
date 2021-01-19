
// 到达火星需要多长时间
package main
import "fmt"
func main() {
	const lightSpeed = 299792  // km/s
	var distance = 5600000000 // km
	fmt.Println(distance/lightSpeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
	const hoursPerDay = 24
	var speed = 100800 // km/h
	var distance2 = 96300000 //km
	// 一次声明多个变量
	//var (
	//	speed = 100800
	//	distance2 = 96300000
	//)
	// var distance2, speed = 96300000, 100800
	fmt.Println(distance2/speed/hoursPerDay, "days")
}
