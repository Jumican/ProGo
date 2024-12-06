package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	
	fmt.Print(a*b*0.5)
}