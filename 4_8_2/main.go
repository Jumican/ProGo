package main

import (
	"fmt"
)
 
func main() {
  var n int 

  for true {

    fmt.Scan(&n)

    if n > 100 {
      break
    } else if n >= 10 {
      fmt.Println(n)
    }

  } 
 
}