package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	fmt.Println(min((a+b+c), (a+c+a+c), (a+a+b+b), (b+b+c+c)))
}