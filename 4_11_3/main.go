package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
    var n string
    fmt.Scan(&n)
    
    n = strings.ReplaceAll(n,"5","")
    n = strings.ReplaceAll(n,"7","")

    var n1 int 
    n1, _ = strconv.Atoi(n)
    fmt.Println(n1)



}