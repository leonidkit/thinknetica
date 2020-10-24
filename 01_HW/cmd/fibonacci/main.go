package main

import (
	"first/pkg/fibonacci"
	"flag"
	"fmt"
)

func main() {
	var nFib int
	flag.IntVar(&nFib, "n", 0, "ordinal number of the Fibonacci number")
	flag.IntVar(&nFib, "number", 0, "ordinal number of the Fibonacci number")
	flag.Parse()

	fmt.Printf("%d'е число Фибоначчи: %d\n", nFib, fibonacci.Fibonacci(nFib)[nFib])
	fmt.Printf("%d'е число Фибоначчи: %d\n", 10, fibonacci.Fibonacci(10)[10])
	fmt.Printf("%d'е число Фибоначчи: %d\n", 15, fibonacci.Fibonacci(15)[15])
}
