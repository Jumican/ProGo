package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a == b && a == c {
		fmt.Println("3")
	} else {
			if a != b && a != c && b != c {
				fmt.Println("0")
			} else{
				fmt.Println("2")
			}
		}
}