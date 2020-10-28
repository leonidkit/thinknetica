package fakebot

import (
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	want := map[string]string{
		"https://habr.ru":         "Главная",
		"https://habr.ru/contact": "Контакты",
	}

	fk := NewBot()
	got, err := fk.Scan("https://example.com", 1)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatal("results is not equal")
	}
}
