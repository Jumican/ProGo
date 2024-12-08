package main

import (
	"fmt"
	"strconv"
)

func main() {
  var n int
  var c string

  fmt.Scan(&n)

  for n > 0 {
      c += strconv.Itoa(n % 10)
      n = n / 10
  }

  fmt.Println(c)
}