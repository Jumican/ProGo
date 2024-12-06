package main

import (
	"fmt"
)

func main() {
	var r float64

	fmt.Scanln(&r)
	
	fmt.Print(r*r*3.14)
}