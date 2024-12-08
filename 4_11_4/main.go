package main

import (
	"fmt"
)

func main() {
    var a, n int
    c := 1
    fmt.Scan(&a, &n)
    
    for i := 0; i < n; i++{
        c *= a
    }

    fmt.Println(c)
}