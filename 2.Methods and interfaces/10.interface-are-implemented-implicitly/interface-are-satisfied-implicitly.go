package main

import "fmt"

// I is interface
type I interface {
	M()
}

// T is struct
type T struct {
	S string
}

// M This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
