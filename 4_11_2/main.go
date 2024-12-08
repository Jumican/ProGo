package main

import (
	"fmt"
)

func main() {
    var n, k int
    s := "YES"

    fmt.Scan(&n)

    for k=10;;k*=10{
        if n%k == n {
            break
        }
    }

    k = k/10

    for n != 0{
        if n%10 != n/k{
            s = "NO"
            break
        }

        n = n%k/10
        k /= 100
    }

    fmt.Println(s)

}