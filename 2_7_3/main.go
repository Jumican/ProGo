package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scanln(&num)
	
	n1 := num%1000/100
	n2 := num%100/10
	n3 := num%10

	fmt.Println(n1+n2+n3)
}