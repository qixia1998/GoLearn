package main

/*
 int sum(int a, int b) {
	 return a + b;
 }
 */
 import "C"
 import "fmt"

 func main() {
	 a := 1
	 b := 1
	 c := C.sum(C.int(a), C.int(b))
	 fmt.Printf("%d + %d = %d\n", a, b, c)
 }