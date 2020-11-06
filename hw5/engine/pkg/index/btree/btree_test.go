package btree

import (
	"testing"
)

type Document struct {
	ID    uint64
	Title string
	URL   string
}

func (d *Document) Ident() uint64 {
	return d.ID
}

func TestSearch(t *testing.T) {
	tr := New()

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
		if got.Ident() != want.Ident() {
			t.Fatalf("error searching an item in tree: want %d, but got %d", want.Ident(), got.Ident())
		}
	}
}
