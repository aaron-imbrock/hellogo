package main

import "fmt"

func man() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)
	// := syntax is shorthand for declaring and initializng a variable.
	//	This syntax is only available inside functions.
	f := "apple"
	fmt.Println(f)
}
