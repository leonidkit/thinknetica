package fakeindex

import "gosearch/pkg/crawler"

type FakeIndex struct{}

func New() *FakeIndex {
	return &FakeIndex{}
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
