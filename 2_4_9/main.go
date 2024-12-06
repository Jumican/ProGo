package main

import (
	"fmt"
)

func main() {
	var rub int
	var cop int
	var n int

	fmt.Scanln(&rub)
	fmt.Scanln(&cop)
	fmt.Scanln(&n)

	r1 := rub*n+cop*n/100
	c1 := cop*n%100

	fmt.Println(r1, c1)
}