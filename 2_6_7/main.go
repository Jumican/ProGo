package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	c := math.Sqrt(a*a+b*b)
	fmt.Println(a + b + c)
}