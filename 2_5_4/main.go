package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scanln(&num)

	n1 := num/100
	n2 := num%100/10
	n3 := num%10
	
	fmt.Print(n3)
	fmt.Print(n2)
	fmt.Print(n1)
}