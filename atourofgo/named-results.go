package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum -x
	return		// A return statement w/o arguments returns the named values. "Naked" return.
}

func main() {
	fmt.Println(split(17))
}
