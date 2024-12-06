package main

import (
	"fmt"
)

func main() {
	var n1 int
	var n2 int
	var n3 int
	var n4 int

	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)
	fmt.Scanln(&n4)

	fmt.Println((n1+n2+n3+n4)*3)
}