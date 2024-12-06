package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scanln(&num)

	fmt.Println(num%1000/100)
}