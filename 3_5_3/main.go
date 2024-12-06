package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scan(&x)

	if x%2 == 1 || x >= 6 && x <= 20{
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}