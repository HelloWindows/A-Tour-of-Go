package main

import "fmt"

func fibonacci() func() int {
	/******* version 1 ********
	index := 0
	last := 0
	result := 0
	ref := 0
	return func() int {
		switch index {
		default:
			ref = result
			result = result + last
			last = ref
		case 0:
			result = 0
		case 1:
			result = 1
		}
		index++
		return result
	}
	*/
	/******* version 2 ********
	x, y, v := 0, 1, 0
	return func() int {
		v, x, y = x, y, x+y
		return v
	}
	*/

	/******* version 3 ********/
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return y - x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
