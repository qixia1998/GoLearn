package main

import (
	_ "GoLearn/save_the_world_with_go/02_the_basics/init_functions/example_02/a"
	"fmt"
)

func init() {
	fmt.Println("Init from my program")
}

func main() {
	fmt.Println("My program")
}
