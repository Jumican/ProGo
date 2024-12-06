package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	fmt.Println(a+b+c, a*b*c)
}