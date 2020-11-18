package btree

import (
	"fmt"
	"gosearch/pkg/crawler"
	"testing"
)

func TestAdd(t *testing.T) {
	tr := New()

	testCases := []struct {
		IsError   bool
		ErrorText string
		Want      *crawler.Document
	}{
		{
			IsError:   false,
			ErrorText: "",
			Want: &crawler.Document{
				uint64(2),
				"",
				"",
			},
		},
		{
			IsError:   false,
			ErrorText: "",
			Want: &crawler.Document{
				uint64(3),
				"",
				"",
			},
		},
		{
			IsError:   true,
			ErrorText: "element already exist",
			Want: &crawler.Document{
				uint64(2),
				"",
				"",
			},
		},
		{
			IsError:   true,
			ErrorText: "element is nil",
			Want:      nil,
		},
	}

	for i, tcase := range testCases {
		err := tr.Add(tcase.Want)
		if tcase.IsError && err != nil {
			if err.Error() != tcase.ErrorText {
				t.Fatalf("[%d] ожидалась ошибка \"%s\", а получена \"%s\"", i, tcase.ErrorText, err.Error())
			}
		}

		if tcase.IsError && err == nil {
			t.Fatalf("[%d] ожидалась ошибка \"%s\", но ошибка не получена", i, tcase.ErrorText)
		}

		if !tcase.IsError && err != nil {
			t.Fatalf("[%d] ожидалось \"%v\", а получена ошибка \"%s\"", i, tcase.Want, err.Error())
		}
	}

	want := "\n\t3\n2\n"
	if fmt.Sprint(tr) != want {
		t.Fatalf("ожидалось \"%v\", а получено \"%s\"", want, tr)
	}
}

func TestSearch(t *testing.T) {
	tr := New()

	for i := 0; i < 3; i++ {
		err := tr.Add(&crawler.Document{ID: uint64(i), Title: "Title", URL: "URL"})
		if err != nil {
			t.Fatalf("ошибка добавления элемента с ID=%d в дерево: %s", i+1, err.Error())
		}
	}

	type TestCase struct {
		IsError   bool
		ErrorText string
		Want      *crawler.Document
	}

	testCases := []struct {
		IsError   bool
		ErrorText string
		Want      *crawler.Document
	}{
		{
			IsError:   false,
			ErrorText: "",
			Want: &crawler.Document{
				uint64(2),
				"",
				"",
			},
		},
		{
			IsError:   true,
			ErrorText: "document not found",
			Want: &crawler.Document{
				uint64(3),
				"",
				"",
			},
		},
		{
			IsError:   true,
			ErrorText: "element is nil",
			Want:      nil,
		},
	}

	for i, tcase := range testCases {
		got, err := tr.Search(tcase.Want)
		if tcase.IsError && err != nil {
			if err.Error() != tcase.ErrorText {
				t.Fatalf("[%d] ожидалась ошибка \"%s\", а получена \"%s\"", i, tcase.ErrorText, err.Error())
			}
		}

		if tcase.IsError && err == nil {
			t.Fatalf("[%d] ожидалась ошибка \"%s\", но ошибка не получена", i, tcase.ErrorText)
		}

		if !tcase.IsError && err != nil {
			t.Fatalf("[%d] ожидалось \"%v\", а получена ошибка \"%s\"", i, tcase.Want, err.Error())
		}

		if !tcase.IsError && err != nil {
			if got.Ident() != tcase.Want.Ident() {
				t.Fatalf("[%d] ошибка поиска элемента в дереве: ожидалось \"%d\", но получили \"%d\"", i, tcase.Want.Ident(), got.Ident())
			}
		}
	}
}
