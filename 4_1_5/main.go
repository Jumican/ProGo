package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	for ; b >= a; b-- {
		fmt.Println(b)
	}
}
