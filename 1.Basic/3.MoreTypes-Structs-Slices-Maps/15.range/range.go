package main

import "fmt"

var pow = []int{0, 1, 2, 3, 4, 5}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}