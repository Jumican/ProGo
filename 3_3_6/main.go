package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64

	fmt.Scan(&a, &b, &c)

	d = b*b - 4*a*c
	if d == 0 {
		fmt.Println(-b/(2*a))
	} else {
		if d > 0 {
			x1 := (-b + math.Sqrt(d))/(2*a)
			x2 := (-b - math.Sqrt(d))/(2*a)
			fmt.Println(math.Min(x1, x2))
			fmt.Println(math.Max(x1, x2))
		}
	}
}