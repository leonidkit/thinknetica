package main

import (
	_ "fmt"
)

func Sum(a int, b int) int {
	return a + b
}

func main() {
	_ = Sum(1, 2)
}
