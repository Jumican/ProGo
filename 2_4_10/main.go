package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scanln(&num)

	h := num/60
	m := num%60

	fmt.Println(num, "мин - это", h, "час", m, "минут.")
}