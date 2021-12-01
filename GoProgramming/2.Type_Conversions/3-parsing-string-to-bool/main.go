package main

import (
	"strconv"
	"fmt"
)

func main() {
	isNew := "j";
	isNewBool, err := strconv.ParseBool(isNew)
	if(err != nil){
		fmt.Println("failed")
	} else {
		if(isNewBool){
			fmt.Println("IsNew")
		} else {
			fmt.Println("Not new")
		}
	}
}