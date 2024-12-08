package main

import (
	"fmt"
)

func main() {
    var n int
    _, _ = fmt.Scan(&n)

    var max, min int
    var number int

    _, _ = fmt.Scan(&max)
    min = max

    for i := 1; i < n; i++ {
        _, _ = fmt.Scan(&number)

        if number > max {
            max = number
        }

        if number < min {
            min = number
        }
    }

    fmt.Println(max-min)
}