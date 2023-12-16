package main

import "fmt"

func main() {
	i := 20
	j := 100 // set initial values

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 5          // set i through the pointer
	fmt.Println(i)  // see the new value of i

	x := &j         // point to j
	*x = *x / 2     // divide j through the pointer
	fmt.Println(*x) // see the new value of j
}
