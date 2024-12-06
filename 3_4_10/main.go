package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	a1 := math.Ceil(a/2)
	b1 := math.Ceil(b/2)
	c1 := math.Ceil(c/2)

	fmt.Println(a1+b1+c1)
}