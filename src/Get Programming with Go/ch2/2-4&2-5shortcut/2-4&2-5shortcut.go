
// 增量并赋值操作符

package main
import "fmt"
func main() {
	var weight = 149.0
	//weight = weight * 0.3783
	weight *= 0.3783
	fmt.Println(weight)
	
	var age = 41
	//age = age + 1
	age += 1
	fmt.Println(age)
	
	var count = 10
	count --
	fmt.Println(count)
	
	var price = 100
	price /= 2
	fmt.Println(price)
}
