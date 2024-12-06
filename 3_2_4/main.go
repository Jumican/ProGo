package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scan(&x)

	if x%10 == x%100/10 || x%10 == x/100 || x/100 == x%100/10 {
		fmt.Println("NO")
	} else {
			fmt.Println("YES")
		}
}