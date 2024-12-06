package main

import (
	"fmt"
)

func main() {
	var x, y string

	fmt.Scan(&x, &y)

	if x == y {
		fmt.Println("Пароль принят")
	} else {
			fmt.Println("Пароль не принят")
	}
}