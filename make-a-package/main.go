package main

import (
	"fmt"

	"example/user/pack/formatter"
	"example/user/pack/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}

