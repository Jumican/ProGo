package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scan(&x)

	if x%10%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}