package main

import (
	"fmt"
)

func main() {
  var n, l, c int

  fmt.Scan(&n)

  for n != 0 {
    if n > l {
      c++
    }

    l = n
    fmt.Scan(&n)
  }

  fmt.Println(c-1)
}