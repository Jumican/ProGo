package main

import (
	"fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    c := n

    for c >= 10 {
        n = c
        c = 0

        for n != 0 {
            c += n%10
            n /= 10
        }
    }
    fmt.Println(c)
}