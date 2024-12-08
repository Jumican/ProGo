package main

import (
	"fmt"
)

func main() {
    var c, sum int
    for i := 6; c != 4; i++ {
        for j := 1; j <= i/2; j++ {
            if i%j == 0 {
                sum += j
            }
        }

        if i == sum {
            fmt.Print(i, ",")
            c++
        }

        sum = 0
    }
}