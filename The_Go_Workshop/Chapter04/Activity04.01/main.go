package main
import "fmt"

func getArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

func main() {
	var arr [10]int
	fmt.Println(getArray(arr))
}