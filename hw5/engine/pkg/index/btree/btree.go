package btree

import (
	"errors"
	"strconv"
)

type Valuer interface {
	ID() uint64
}

type Tree struct {
	Left  *Tree
	Value Valuer
	Right *Tree
}

func NewTree() *Tree {
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
	if d == nil {
		return errors.New("element is nil")
	}
	if t == nil {
		return errors.New("tree is nil")
	}
	if t.Value == nil {
		t.Value = d
		return nil
	}

	if d.(Valuer).ID() == t.Value.ID() {
		return errors.New("element already exist")
	}
	if d.(Valuer).ID() < t.Value.ID() {
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
	if d.(Valuer).ID() > t.Value.ID() {
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
	if d == nil {
		return nil, errors.New("element is nil")
	}
	if t == nil {
		return nil, errors.New("document not found")
	}

	if t.Value.ID() == d.(Valuer).ID() {
		return t.Value, nil
	}

	if d.(Valuer).ID() < t.Value.ID() {
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
	res += strconv.FormatUint(t.Value.ID(), 10) + "\n"
	res += t.Left.Print(depth + 1)
	return res
}

func (t *Tree) String() string {
	return "\n" + t.Print(0)
}
