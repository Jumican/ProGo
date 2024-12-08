package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    c := 1
    for i := 2; i <= n; i++ {
        if i%2 == 0{
            c *= i
        }
    }
    fmt.Println(c)
}