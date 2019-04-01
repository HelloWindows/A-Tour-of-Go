package main

import "fmt"

func main() {
	primes := []int{0, 1, 2, 3, 4, 5}
	s := primes[1:4]
	fmt.Println(s)
	// The default is zero for the low bound
	fmt.Println("primes[:3] =", primes[:3])
	// The default is the length of the slice for the high bound
	fmt.Println("primes[4:] =", primes[4:])
	// The default is zero for the low bound and the lenght of the slice for high bound
	fmt.Println("primed[:]", primes[:])
}
