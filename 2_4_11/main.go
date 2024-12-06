package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scanln(&num)

	next := num+1
	prev := num-1

	fmt.Println("Следующее за числом", num, "число:", next)
	fmt.Println("Для числа", num, "предыдущее число:", prev)
}