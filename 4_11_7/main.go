package main

import (
	"fmt"
	"os"
)

func main() {
    var a, b int
    fmt.Scan(&a, &b)

    for ;a <= b;b-- {
        if b%7 == 0 {
            fmt.Println(b)
            os.Exit(0)
        }
    }
    fmt.Println("NO")
}