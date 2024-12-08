package main

import (
	"fmt"
)

func main() {
    var n, number, max int
    s := "NO"

    _, _ = fmt.Scan(&n)

    _, _ = fmt.Scan(&max)

    if max < 30 {
        s = "YES"
    }

    for i := 1; i < n; i++ {
        _, _ = fmt.Scan(&number)

        if number > max {
            max = number
        }

        if number < 30 {
            s = "YES"
        }
    }

    fmt.Println(max, s)
}