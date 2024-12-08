package main

import (
	"fmt"
)

func main() {
  var a, b, k, c int

  fmt.Scan(&a, &b, &k)

  if k <= 1 {

    for ;a <= b; a++ {
      fmt.Print(a, " ")
    }

  } else if k == 2{

      if a == 1 {
        a++
      }

      for ;a <= b; a++ {
        fmt.Print(a, " ")
      }

  } else {

      for ;a <= b; a++ {

        for i := 2; i <= a/2; i++ {

          if a%int(i) == 0 {

            c++
            if c >= k-2 {

              fmt.Print(a, " ")
              i = a

            }

          }

        }

        c = 0

      }

    }
    
}