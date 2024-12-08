package main

import (
	"fmt"
)

func main() {
    var c int

    for i := 100; i < 1000; i++ {
        if i%10 + i/10%10 + i/100 == 8 {
            c++
        }
    }

    fmt.Println(c)

}