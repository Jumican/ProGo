package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Scan(&x)

	if x <= -3 || x >= 7 {
		fmt.Println("Принадлежит")
	} else {
			fmt.Println("Не принадлежит")
		}
}