package main

import (
	"fmt"
)

func main() {
	var a float64

	fmt.Scanln(&a)
	var b = int(a)
	fmt.Print(a-float64(b))
}