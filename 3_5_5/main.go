package main

import (
	"fmt"
)

func main() {
	var x float64

	fmt.Scan(&x)

	if x <= 3{
		fmt.Println("начинающий")
	} else if x <= 7{
		fmt.Println("младший разработчик")
	} else if x <= 15{
		fmt.Println("средний разработчик")
	} else {
		fmt.Println("старший разработчик")
	}

}