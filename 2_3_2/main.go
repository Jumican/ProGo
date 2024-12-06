package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	book := scanner.Text()
  fmt.Print(book, " - лучшая книга!")
}