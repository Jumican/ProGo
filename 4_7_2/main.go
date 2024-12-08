package main

import (
	"fmt"
)

func main() {
  var a, c int

  fmt.Scan(&a)

  for j := 1; j <= a; j++ {
    for i := j; i != 0; i /= 10 {
      if i%10 == 7 {
        c++
      }
    }
  }

  fmt.Println(c)

}