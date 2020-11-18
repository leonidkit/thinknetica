package btree

import (
	"errors"
	"reflect"
	"strconv"
)

type Valuer interface {
	Ident() uint64
}

type Tree struct {
	Left  *Tree
	Value Valuer
	Right *Tree
}

func New() *Tree {
	return &Tree{
		Left:  nil,
		Value: nil,
		Right: nil,
	}
}

func (t *Tree) Add(d Valuer) error {
	if _, ok := d.(Valuer); !ok {
		return errors.New("the passed argument does not implement the interface Valuer")
	}
	if isNilFixed(d) {
		return errors.New("element is nil")
	}
	if t == nil {
		return errors.New("tree is nil")
	}
	if t.Value == nil {
		t.Value = d
		return nil
	}

	if d.(Valuer).Ident() == t.Value.Ident() {
		return errors.New("element already exist")
	}
	if d.(Valuer).Ident() < t.Value.Ident() {
		if t.Left == nil {
			t.Left = &Tree{
				Left:  nil,
				Value: d,
				Right: nil,
			}
			return nil
		}
		t.Left.Add(d)
	}
	if d.(Valuer).Ident() > t.Value.Ident() {
		if t.Right == nil {
			t.Right = &Tree{
				Left:  nil,
				Value: d,
				Right: nil,
			}
			return nil
		}
		t.Right.Add(d)
	}
	return nil
}

func (t *Tree) Search(d Valuer) (Valuer, error) {
	if _, ok := d.(Valuer); !ok {
		return nil, errors.New("the passed argument does not implement the interface Valuer")
	}
	if reflect.ValueOf(d).IsNil() {
		return nil, errors.New("element is nil")
	}
	if t == nil {
		return nil, errors.New("document not found")
	}

	if t.Value.Ident() == d.(Valuer).Ident() {
		return t.Value, nil
	}

	if d.(Valuer).Ident() < t.Value.Ident() {
		return t.Left.Search(d)
	} else {
		return t.Right.Search(d)
	}
}

func (t *Tree) Print(depth int) (res string) {
	if t == nil {
		return
	}
	res += t.Right.Print(depth + 1)
	for i := 0; i < depth; i++ {
		res += "\t"
	}
	res += strconv.FormatUint(t.Value.Ident(), 10) + "\n"
	res += t.Left.Print(depth + 1)
	return res
}

func (t *Tree) String() string {
	return "\n" + t.Print(0)
}

func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		//use of IsNil method
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
