package main

import "fmt"

func main() {
    var n, x int
    fmt.Scan(&n)
    count := 0
    for i := 1; i <= n; i++ {
				fmt.Scan(&x)
        if x == 0 {
            count++
        }
    }

    if count > 0 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}