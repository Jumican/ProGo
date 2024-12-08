package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scan(&x)

	for i := 1; i <= x/2; i++ {
		if x%i == 0 {
			fmt.Println(i)
		}
	}
    fmt.Println(x)
}
