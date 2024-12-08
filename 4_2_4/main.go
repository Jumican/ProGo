package main

import "fmt"

func main() {
    var n, x int
    fmt.Scan(&n)
    count := 0
    for i := 1; i <= n; i++ {
				fmt.Scan(&x)
        if x%10 == 0 {
            count++
        }
    }

    fmt.Println(count)
}