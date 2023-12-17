package main

import (
	"fmt"
	"math"
)

// #3 As defined the compute function takes a function (`fn`) as a paramter, expecting that
//    function to accept two `float64` params and to return a `float64`. It then calls this function
//    (`fn`) with the args `3` and `4` and returns the result.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func main() {
	// #1
	// anonymous function, that takes two `float64` params(x and y)
	// and returns the square root of the sum of their squares.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// #2
	// Calls `hypot` function w/ args 5 and 12 and prints the result
	fmt.Println(hypot(5, 12))
	
	// #4
	// the `compute` function take the `hypot` function as a paramter, 
	// it then calls this function with the args `3` and `4` and returns the result.
	fmt.Println(compute(hypot))
	// #5
	// the `compute` function takes `math.Pow` as a paramter and calls it with
	// the args `3` and `4`, effectively raising 3 to the power of 4.
	fmt.Println(compute(math.Pow))
}

