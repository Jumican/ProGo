package main

import (
	"fmt"
)

func main() {
	var n, x, c int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		c += x
	}
  fmt.Println(c)
}