package main

import (
	"fmt"
)

func main() {
	var age int
	var gender string

	fmt.Scan(&age, &gender)

	if gender == "m" && age >= 12 && age <= 18{
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}