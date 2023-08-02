package main

import (
	"GoLearn/FunctionalProgramming/ch04/Examples/HotdogShop/PureHotdogShop/pkg"
	"fmt"
)

func main() {
	myCard := pkg.NewCreditCard(1000)
	hotdog, creditFunc := pkg.OrderHotdog(myCard, pkg.Charge)
	fmt.Printf("%+v\n", hotdog)
	newCard, err := creditFunc()
	if err != nil {
		panic("User has no credit")
	}
	myCard = newCard
	fmt.Printf("%+v\n", myCard)
}
