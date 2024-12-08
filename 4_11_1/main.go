package main

import (
	"fmt"
)

func main() {
    var n, num, c int

    fmt.Scan(&n)

    for i := 1; i <= n; i++{
        fmt.Scan(&num)

        if num == 0 {
            c++
        }

    }

    fmt.Println(c)
}