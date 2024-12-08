package main

import (
	"fmt"
)

func main() {
  var n, l, c int

  fmt.Scan(&n)

  for n != 0 {
    if n > 0 && l < 0 || n < 0 && l > 0 {
      c++
    }

    l = n
    fmt.Scan(&n)
  }

  fmt.Println(c)
}