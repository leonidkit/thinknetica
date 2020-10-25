package main

import (
	"fmt"
)

func main() {
	a := 'A'
	b := "A"

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	switch a {
	case 'A':
		{
			print("trw")
			print("\n")
		}
	}

	lsice := []int{1, 2, 3, 4, 5}
	slice := lsice[2:3]
	fmt.Println(slice)

	slice[0] = 110

	fmt.Println(lsice)
}
