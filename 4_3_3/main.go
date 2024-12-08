package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    c := 1
    for ; a <= b; a++ {
        if a%10 == 7 {
            c *= a
        }
    }
    fmt.Println(c)
}