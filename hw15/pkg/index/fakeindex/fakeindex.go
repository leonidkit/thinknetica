package fakeindex

import (
	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
)

type FakeIndex struct{}

func New() *FakeIndex {
	return &FakeIndex{}
}

func (f *FakeIndex) Recieve() index.Index {
	return index.Index{"как": []uint64{1123123, 12432343, 1242544}}
}

func (f *FakeIndex) Find(request string) ([]crawler.Document, error) {
	return []crawler.Document{
		crawler.Document{
			ID:    uint64(1),
			Title: "Как использовать git?",
			URL:   "http://localhost",
		},
		crawler.Document{
			ID:    uint64(2),
			Title: "Прикладное применение подорожника как лекарство",
			URL:   "http://localhost",
		},
		crawler.Document{
			ID:    uint64(3),
			Title: "Криптовалюта как будущее финансовой системы?",
			URL:   "http://localhost",
		},
	}, nil
}
