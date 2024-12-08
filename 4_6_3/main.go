package main

import (
	"fmt"
)

func main() {
  var n float64
  var sum, c float64

  fmt.Scan(&n)

  for n != 0 {
    c++
    sum += n
    fmt.Scan(&n)
  }

  fmt.Println(sum/c)
}