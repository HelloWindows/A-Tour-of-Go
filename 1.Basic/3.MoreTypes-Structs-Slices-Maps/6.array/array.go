package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(primes)

	for i := 0; i < len(primes); i++ {
		fmt.Printf("primes[%d] = % d\n", i, primes[i])
	}
}
