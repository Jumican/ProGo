package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	for a += a%2*a%2; a <= b; a = a + 2 {
		fmt.Println(a)
	}
}