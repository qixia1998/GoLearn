package main

import "fmt"

// 短路：评估go中的条件

func main() {

	if raining() || snowing() {
		fmt.Println("Stay indoors!")
	}

	if !raining() || snowing() {
		fmt.Println("Let's ski!")
	}

	if !raining() && !snowing() {
		fmt.Println("Let's go outdoors!")
	}

	if raining() && snowing() {
		fmt.Println("It's going to be really cold!")
	}

}

func raining() bool {
    fmt.Println("Check if it is raining now...")
    return true
}

func snowing() bool {
    fmt.Println("Check if it is snowing now...")
    return true
}