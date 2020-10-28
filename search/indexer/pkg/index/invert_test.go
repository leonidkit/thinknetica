package index

import (
	"strconv"
	"testing"
)

func BenchmarkNewIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var data = map[string]string{
			"http://habr.ru/main":    "Главная",
			"http://habr.ru/contact": "Контакты",
			"http://habr.ru/comment": "Комментарии",
			"http://habr.ru/lol":     "Лол",
		}
		NewIndex(data)
	}
}

func TestBinarySearch(t *testing.T) {
	testDocs := []Document{}
	for i := 0; i < 10000; i++ {
		testDocs = append(testDocs, Document{
			uint64(i), "Title" + strconv.Itoa(i), "URL" + strconv.Itoa(i),
		})
	}

	for i := 0; i < 1000; i++ {
		_, err := BinarySearch(uint64(i+1), testDocs)
		if err != nil {
			t.Fatalf("%d must be found in %d case", i, i)
		}
	}
}
