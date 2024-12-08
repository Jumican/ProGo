package main

import (
	"fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  c := 0

  for n != 0 {
      if n < 0 {
        c--
      } else {
        c++
      }

      fmt.Scan(&n)
  }

  fmt.Println(c)
}