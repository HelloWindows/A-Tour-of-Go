package main

import "fmt"

// I is interface
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// This is a run-time error because there is no type inside
	// the interface tuple to indicate which concrete method to call
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
