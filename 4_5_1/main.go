package main

import (
	"fmt"
)

func main() {
    var n, d int
    fmt.Scan(&n)
    с := 0
    for n > 0 {
        d = n % 10
        if d == 4{
          с += 1
        }
        n = n / 10
    }
    fmt.Println(с)
}