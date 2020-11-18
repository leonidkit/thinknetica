package hw11

import (
	"fmt"
	"io"
)

type Human interface {
	Age() int
}

type Employee struct {
	age int
}

type Customer struct {
	age int
}

func (e *Employee) Age() int {
	return e.age
}

func (c *Customer) Age() int {
	return c.age
}

// Возвращает максимальный возраст из передаваемых аргументов типа Human
func MaxAge(people ...Human) int {
	max := 0
	for _, h := range people {
		if h.Age() > max {
			max = h.Age()
		}
	}
	return max
}

// Возвращает аргумент с максимальным возрастом
func MaxAgeHuman(people ...interface{}) interface{} {
	max := struct {
		age  int
		item interface{}
	}{}
	for _, h := range people {
		switch i := h.(type) {
		case *Employee:
			if i.age > max.age {
				max.age = i.age
				max.item = i
			}
			break
		case *Customer:
			if i.age > max.age {
				max.age = i.age
				max.item = i
			}
			break
		case Employee:
			if i.age > max.age {
				max.age = i.age
				max.item = i
			}
			break
		case Customer:
			if i.age > max.age {
				max.age = i.age
				max.item = i
			}
			break
		}

	}
	return max.item
}

// Производит запись аргументов типа string в out
func Print(out io.Writer, items ...interface{}) {
	for _, i := range items {
		switch s := i.(type) {
		case string:
			_, err := out.Write([]byte(s))
			if err != nil {
				fmt.Errorf("printing item %v error: ", err.Error())
				return
			}
			continue
		default:
			continue
		}
	}
}
