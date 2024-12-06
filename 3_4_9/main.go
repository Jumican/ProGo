package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a+b+c - max(a, b, c)*2 > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}