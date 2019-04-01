package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

// Abs is abs
func Abs(v vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale is scale
func Scale(v *vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
