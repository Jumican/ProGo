package main

import (
	"fmt"
	"math"
)

func main() {
	var k2, k3, k5, k6 float64

	fmt.Scan(&k2, &k3, &k5, &k6)
	m := math.Min(k2, k5)
	m = math.Min(m, k6)
	k2 = k2-m
	c := m*256
	m = math.Min(k2, k3)
	c += m*32
	fmt.Println(c)
	
}