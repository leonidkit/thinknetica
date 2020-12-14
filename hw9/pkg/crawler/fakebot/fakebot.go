package fakebot

import (
	"gosearch/pkg/crawler"
)

type Scan struct{}

func New() *Scan {
	return &Scan{}
}

func (l *Scan) Scan(url string, depth int) ([]crawler.Document, error) {
	data := []crawler.Document{
		{
			URL:   "https://habr.ru",
			Title: "Главная",
		},
		{
			URL:   "https://habr.ru/contact",
			Title: "Контакты",
		},
	}

	return data, nil
}
