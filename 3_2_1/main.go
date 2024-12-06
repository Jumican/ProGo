package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scan(&x)

	if x >= 100 && x < 1000 {
		fmt.Println("YES")
	} else {
			fmt.Println("NO")
		}
}