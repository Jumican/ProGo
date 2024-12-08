package main

import (
	"fmt"
)

func main() {
    var y int
    fmt.Scan(&y)

    for y++;;y++{
        if  y%10 != y%100/10 && y%10 != y%1000/100 && y%10 != y/1000 &&
            y%100/10 != y%1000/100 && y%100/10 != y/1000 && y%1000/100 != y/1000 {
            
            fmt.Println(y)
            break

        }
    }
}