package engine

import (
	"gosearch/pkg/index/fakeindex"
	"testing"
)

func TestSearch(t *testing.T) {
	idx := fakeindex.New()

	engine := New(idx)

	want := "как"
	_, err := engine.Search(want)
	if err != nil {
		t.Fatalf("ожидался результат, а получена ошибка: %s", err.Error())
	}

	want = ""
	wantErr := "пустой запрос"
	_, err = engine.Search(want)
	if err != nil {
		if err.Error() != wantErr {
			t.Fatalf("ожидалась ошибка %s, а получена: %s", wantErr, err.Error())
		}
	}
}
