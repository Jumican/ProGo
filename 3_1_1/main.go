package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64

	fmt.Scanln(&x)
	fmt.Scanln(&y)

	fmt.Println(math.Max(x, y))
}