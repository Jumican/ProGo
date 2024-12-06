package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scan(&x)

	if x%10 + x%100/10 + x%1000/100 == x%10000/1000 + x%100000/10000 + x/100000 {
		fmt.Println("YES")
	} else {
			fmt.Println("NO")
		}
}