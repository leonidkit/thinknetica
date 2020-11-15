package invert

import (
	"strconv"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testDocs := []Document{}
	for i := 0; i < 10; i++ {
		testDocs = append(testDocs, Document{
			uint64(i), "Title" + strconv.Itoa(i), "URL" + strconv.Itoa(i),
		})
	}

	want := uint64(6)
	got, err := BinarySearch(want, testDocs)
	if err != nil {
		t.Fatalf("must be fount %d, but error", want)
	} else {
		if got != want {
			t.Fatalf("must be fount %d, but found %d", want, got)
		}
	}

	want = uint64(11)
	wantErr := "nothing found"
	_, err = BinarySearch(want, testDocs)
	if err == nil {
		t.Fatalf("must be error but no")
	} else {
		if err.Error() != wantErr {
			t.Fatalf("must return message: \"%s\", but got \"%s\"", wantErr, err.Error())
		}
	}
}
