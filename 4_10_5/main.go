package main

import (
	"fmt"
)

func main() {
    var number, max0, max1, temp int

    fmt.Scan(&max0, &max1)

    temp = min(max0, max1)
    max0 = max(max0, max1)
    max1 = temp

    for {
        fmt.Scan(&number)

        if number == 0 {
            break
        }

        if number > max0 {
            max1 = max0
            max0 = number
        } else if number > max1 {
            max1 = number
        }

    }

    fmt.Println(max1)
}