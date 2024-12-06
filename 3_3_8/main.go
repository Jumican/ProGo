package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Scan(&a)

	switch a {
		case 1: fmt.Println("Зима")
		case 2: fmt.Println("Зима") 
		case 3: fmt.Println("Весна")
		case 4: fmt.Println("Весна")
		case 5: fmt.Println("Весна")
		case 6: fmt.Println("Лето")
		case 7: fmt.Println("Лето")
		case 8: fmt.Println("Лето")
		case 9: fmt.Println("Осень")
		case 10: fmt.Println("Осень")
		case 11: fmt.Println("Осень")
		case 12: fmt.Println("Зима")		
	}
}