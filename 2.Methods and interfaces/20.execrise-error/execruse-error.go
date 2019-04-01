package main

import (
	"fmt"
	"math"
)

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

const offset = 1e-8

// Sqrt is func
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errNegativeSqrt(x)
	}
	t, z := 0., 1.
	for {
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) < offset {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
