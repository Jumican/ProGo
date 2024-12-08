package main

import (
	"fmt"
)

func main() {
  var a, c, a1, c1 int

  fmt.Scan(&a)

  for j := 1; j <= a; j++ {

    for i := j; i != 0; i /= 10 {
        c += i%10

    }

    if c > c1 {
      c1 = c
      a1 = j
    }

    c = 0

  }

  fmt.Println(a1, c1)

}