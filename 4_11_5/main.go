package main

import (
	"fmt"
)

func main() {
    var n, min int
    fmt.Scan(&n)
    fmt.Scan(&min)
    cnt := 1

    for i := 1; i < n; i++ {
        var k int
        fmt.Scan(&k)
        if k < min {
            min = k
            cnt = 1
        } else {
            if k == min {
                cnt++
            }
        }
    }

    fmt.Println(cnt)
}