package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    c := 0
    c1 := 1
    for c1 < n{
        c += 1
        c1 *= 2
    }
    fmt.Println(c)
}