package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    c := 0
    for n%3 == 0{
        n /= 3
        c++
    }
    fmt.Println(c)
}