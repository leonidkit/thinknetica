// Package spider реализует сканер содержимого веб-сайтов.

// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.

package spider

import (
	"testing"
)

func TestScanSite(t *testing.T) {
	const url = "https://habr.com"
	const depth = 1

	sp := New()
	data, err := sp.Scan(url, depth)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range data {
		t.Logf("%s -> %s\n", v.Title, v.URL)
	}
}
