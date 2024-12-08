package main

import (
	"fmt"
)

func main() {
  var a int

  fmt.Scan(&a)

  for j := 2; j <= a; j++ {

    if a%j == 0 {
      fmt.Print(j, " ")
      a /= j
      j--
    }

  }

}