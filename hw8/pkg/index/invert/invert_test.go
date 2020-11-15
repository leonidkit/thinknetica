package invert

import (
	"gosearch/pkg/crawler"
	"strconv"
	"testing"
)

var data = []crawler.Document{
	crawler.Document{
		ID:    uint64(1),
		Title: "Как использовать git?",
		URL:   "http://localhost",
	},
	crawler.Document{
		ID:    uint64(2),
		Title: "Прикладное применение подорожника",
		URL:   "http://localhost",
	},
	crawler.Document{
		ID:    uint64(3),
		Title: "Криптовалюта - будущее финансовой системы?",
		URL:   "http://localhost",
	},
}

func TestFind(t *testing.T) {
	index := NewIndexTree(data)

	want := data[1]
	got, err := index.Find("прикладное")
	if err != nil {
		t.Fatalf("ожидалось %v, получена ошибка %s", want, err.Error())
	}
	if len(got) != 1 {
		t.Fatalf("ожидалось %v, получено %v", want, got)
	}
	if got[0].URL != want.URL {
		t.Fatalf("ожидалось %v, получено %v", want, got)
	}
}

func TestBinarySearch(t *testing.T) {
	testDocs := []crawler.Document{}
	for i := 0; i < 10; i++ {
		testDocs = append(testDocs, crawler.Document{
			uint64(i), "Title" + strconv.Itoa(i), "URL" + strconv.Itoa(i),
		})
	}

	want := uint64(9)
	got, err := binarySearch(want, testDocs)
	if err != nil {
		t.Fatalf("must be fount %d, but error", want)
	} else {
		if got != want {
			t.Fatalf("must be fount %d, but found %d", want, got)
		}
	}

	want = uint64(11)
	wantErr := "nothing found"
	_, err = binarySearch(want, testDocs)
	if err == nil {
		t.Fatalf("must be error but no")
	} else {
		if err.Error() != wantErr {
			t.Fatalf("must return message: \"%s\", but got \"%s\"", wantErr, err.Error())
		}
	}
}
