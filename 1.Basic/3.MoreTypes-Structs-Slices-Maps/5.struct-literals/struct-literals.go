package main

import "fmt"

type vertex struct {
	X, Y int
}

var (
	v1 = vertex{1, 2}  // has type vertex
	v2 = vertex{X: 1}  // T:0 is implicit
	v3 = vertex{}      // X:0 and Y:0
	p  = &vertex{1, 2} // has type *vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
