package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	_ = scanner.Scan()
	di := scanner.Text()

	_ = scanner.Scan()
	s1 := scanner.Text()

	_ = scanner.Scan()
	s2 := scanner.Text()

	_ = scanner.Scan()
	s3 := scanner.Text()

  fmt.Println(s1 + di + s2 + di + s3)
}