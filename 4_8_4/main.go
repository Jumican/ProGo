package main

import (
	"fmt"
)
 
func main() {
  var s, p string
  var c int 

  fmt.Scan(&s)

  for true {
    fmt.Scan(&p)

    c++
    
    if s == p {
      fmt.Println(c)
      break
    }
  } 
}