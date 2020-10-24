package main

import (
	"errors"
	"fmt"
)

type Pointer interface {
	String() string
}

type Point struct {
	X, Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("{%d, %d}", p.X, p.Y)
}

func main() {
	// var pI Pointer
	// var p = Point{1, 2}
	// pI = &p
	// fmt.Println(pI)
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	tmp, ok := <-ch
	fmt.Println(ok)

	close(ch)
	tmp, ok = <-ch
	fmt.Println(ok)
	tmp = <-ch

	_ = tmp
}

func myErrors() {
	// первый способ создания ошибки
	var err = errors.New("sd")

	// второй сопособ
	err = fmt.Errorf("Some text")

	_ = err

	// возвращение цепочек ошибок в GO 1.13
}
