package main

import (
	"fmt"
)

func main() {
	var x, c int

	fmt.Scan(&x)

	for i := 1; i <= x/2; i++ {
		if x%i == 0 {
			c++
		}
	}
    fmt.Println(c+1)
}