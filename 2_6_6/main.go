package main

import (
	"fmt"
)

func main() {
	var a float64

	fmt.Scanln(&a)
	fmt.Print(a/8192)
}