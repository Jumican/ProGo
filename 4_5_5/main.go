package main

import (
	"fmt"
)

func main() {
  var n, c int

  fmt.Scan(&n)

  for n > 0 {
    if n%2 == 1 {
      c += 1
    }
    n /= 2
  }

  fmt.Println(c)
}