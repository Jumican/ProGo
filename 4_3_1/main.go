package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    count := 1
    for i := 1; i <= n; i++ {
        count *= i
    }
    fmt.Println(count)
}