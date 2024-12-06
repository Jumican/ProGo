package main

import (
	"fmt"
)

func main() {
	var k, m int

	fmt.Scan(&k, &m)
	d := k/m
	
	if k%m == 0 {
		fmt.Println(d)
	} else {
		fmt.Println(d+1)
	}
}