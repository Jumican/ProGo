package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)

    for a != b{
        a = max(a, b) - min(a, b)
        b = max(a, b) - min(a, b)
    }

    fmt.Println(a)
}