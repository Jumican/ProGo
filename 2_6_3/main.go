package main

import (
	"fmt"
)

func main() {
	var f float64

	fmt.Scanln(&f)
	
	fmt.Print((f-32)*5/9)
}