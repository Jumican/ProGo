package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanln(&num)
	num = num*num*num
	fmt.Println(num*num)
}