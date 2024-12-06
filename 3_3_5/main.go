package main

import (
	"fmt"
	"os"
)

func main() {
	var n1, n2, sum float64
	var op string

	fmt.Scan(&n1, &n2, &op)

	if op == "+" {
		sum = n1 + n2
	} else {
			if op == "-" {
				sum = n1 - n2
			} else {
					if op == "*" {
						sum = n1 * n2
					} else {
							if op == "/" {
								if n2 != 0 {
									sum = n1 / n2
								} else{
										fmt.Println("На ноль делить нельзя!")
										os.Exit(0)
									}
							} else {
									fmt.Println("Неверная операция")
									os.Exit(0)
								}
						}
				}
		}
	fmt.Println(sum)
}