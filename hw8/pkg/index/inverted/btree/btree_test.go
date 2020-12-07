package btree

import (
	"fmt"
	"gosearch/pkg/crawler"
	"testing"
)

func TestAdd(t *testing.T) {
	tr := New()

	testCases := []struct {
		wantErr   bool
		ErrorText string
		Want      *crawler.Document
	}{
		{
			wantErr:   false,
			ErrorText: "",
			Want: &crawler.Document{
				uint64(2),
				"",
				"",
			},
		},
		{
			wantErr:   false,
			ErrorText: "",
			Want: &crawler.Document{
				uint64(3),
				"",
				"",
			},
		},
		{
			wantErr:   true,
			ErrorText: "element already exist",
			Want: &crawler.Document{
				uint64(2),
				"",
				"",
			},
		},
		{
			wantErr:   true,
			ErrorText: "element is nil",
			Want:      nil,
		},
	}

	for i, tt := range testCases {
		err := tr.Add(tt.Want)
		if tt.wantErr && err != nil {
			if err.Error() != tt.ErrorText {
				t.Fatalf("[%d] ожидалась ошибка \"%s\", а получена \"%s\"", i, tt.ErrorText, err.Error())
			}
		}

		if tt.wantErr && err == nil {
			t.Fatalf("[%d] ожидалась ошибка \"%s\", но ошибка не получена", i, tt.ErrorText)
		}

		if !tt.wantErr && err != nil {
			t.Fatalf("[%d] ожидалось \"%v\", а получена ошибка \"%s\"", i, tt.Want, err.Error())
		}
	}

	want := "\n\t3\n2\n"
	if fmt.Sprint(tr) != want {
		t.Fatalf("ожидалось \"%v\", а получено \"%s\"", want, tr)
	}
}

func TestSearch(t *testing.T) {
	tree := New()

	for i := 0; i < 3; i++ {
		err := tree.Add(&crawler.Document{ID: uint64(i), Title: "Title", URL: "URL"})
		if err != nil {
			t.Fatalf("ошибка добавления элемента с ID=%d в дерево: %s", i+1, err.Error())
		}
	}

	tests := []struct {
		wantErr   bool
		ErrorText string
		Want      *crawler.Document
	}{
		{
			wantErr:   false,
			ErrorText: "",
			Want: &crawler.Document{
				uint64(2),
				"",
				"",
			},
		},
		{
			wantErr:   true,
			ErrorText: "document not found",
			Want: &crawler.Document{
				uint64(3),
				"",
				"",
			},
		},
		{
			wantErr:   true,
			ErrorText: "element is nil",
			Want:      nil,
		},
	}

	for i, tt := range tests {
		got, err := tree.Search(tt.Want)
		if tt.wantErr && err != nil {
			if err.Error() != tt.ErrorText {
				t.Fatalf("[%d] ожидалась ошибка \"%s\", а получена \"%s\"", i, tt.ErrorText, err.Error())
			}
		}

		if tt.wantErr && err == nil {
			t.Fatalf("[%d] ожидалась ошибка \"%s\", но ошибка не получена", i, tt.ErrorText)
		}

		if !tt.wantErr && err != nil {
			t.Fatalf("[%d] ожидалось \"%v\", а получена ошибка \"%s\"", i, tt.Want, err.Error())
		}

		if !tt.wantErr && err != nil {
			if got.Ident() != tt.Want.Ident() {
				t.Fatalf("[%d] ошибка поиска элемента в дереве: ожидалось \"%d\", но получили \"%d\"", i, tt.Want.Ident(), got.Ident())
			}
		}
	}
}
