package main

import (
	"fmt"
	"math"
)

func main() {
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64

	fmt.Scanln(&x1)
	fmt.Scanln(&y1)
	fmt.Scanln(&x2)
	fmt.Scanln(&y2)

	c := math.Abs(x1-x2)+math.Abs(y1-y2)
	fmt.Println(c)
}