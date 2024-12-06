package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scanln(&num)
	n1 := num%10
	n2 := num%100/10
	n3 := num/100

	fmt.Println(n1+n2+n3)
}