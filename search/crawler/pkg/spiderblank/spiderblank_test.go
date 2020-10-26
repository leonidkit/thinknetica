package spiderblank

import "testing"

func AssertEqual(a, b map[string]string) bool {
	for k, v := range a {
		if c, ok := b[k]; ok {
			if c != v {
				return false
			}
		}
	}
	return true
}

func TestScan(t *testing.T) {
	want := map[string]string{
		"https://habr.ru":         "Главная",
		"https://habr.ru/contact": "Контакты",
	}

	got, err := Scan("https://example.com", 1)
	if err != nil {
		t.Fatal(err)
	}

	if !AssertEqual(want, got) {
		t.Fatal("results is not equal")
	}
}
