package main

import (
	"fmt"
)

func main() {
  var n int
  c := 0

  fmt.Scan(&n)
  num := n

  for n > 0 {
      c += n % 10
      n = n / 10
  }

  if num%c == 0 {
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }
}