package example_01_test

import "GoLearn/save_the_world_with_go/11_modules/godoc/example_01"

func ExampleMsg_Send() {
	m := example_01.Msg{"Hello"}
	m.Send("John")
	// Output:
	// Send Hello to John
}

func ExampleMsg_Receive() {
	m := example_01.Msg{"Hello"}
	m.Receive("John")
	// Output:
	// Received Hello from John
}
