package main

import (
	"fmt"
	"math"
)

func oSqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Sqrt(x float64) float64 {
	z := float64(1)
	n := 0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
		if z*z > x {
			z -= 0.00001
		}
	}
	fmt.Printf("We took %v times to guess %v\n", n, z)
	return z
}

func main() {
	fmt.Println("My Sqrt is: ", Sqrt(2))
	fmt.Println("Go's Sqrt is: ", oSqrt(2))
}
