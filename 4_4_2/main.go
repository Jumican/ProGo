package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    c := 1
    for c <= n{
        fmt.Println(c)
        c *= 2 
    }
}