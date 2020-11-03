package btree

import (
	"testing"
)

type Document struct {
	Id    uint64
	Title string
	URL   string
}

func (d *Document) ID() uint64 {
	return d.Id
}

func TestSearch(t *testing.T) {
	tr := NewTree()

	for i := 0; i < 500; i++ {
		err := tr.Add(&Document{Id: uint64(i), Title: "Title", URL: "URL"})
		if err != nil {
			t.Fatalf("error adding an item to tree: %d - %s", i+1, err.Error())
		}
	}

	want := &Document{
		uint64(444),
		"",
		"",
	}
	got, err := tr.Search(want)
	if err != nil {
		t.Fatalf("error searching an item in tree: %s", err.Error())
	} else {
		if got.ID() != want.ID() {
			t.Fatalf("error searching an item in tree: want %d, but got %d", want.ID(), got.ID())
		}
	}
}
