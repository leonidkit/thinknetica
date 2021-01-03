package fakeindex

import (
	"fmt"
	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
	"math/rand"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

type FakeIndex struct {
	index index.Index
	docs  []crawler.Document
}

func New(data []crawler.Document) *FakeIndex {
	return &FakeIndex{
		docs: data,
	}
}

func (f *FakeIndex) Add(url string, title string) error {
	f.docs = append(f.docs, crawler.Document{
		ID:    r.Uint64(),
		Title: title,
		URL:   url,
	})
	return nil
}

func (f *FakeIndex) Delete(ID uint64) error {
	var isDeleted bool = false

	for i, doc := range f.docs {
		if doc.ID == ID {
			isDeleted = true
			f.docs[i] = f.docs[len(f.docs)-1]
			f.docs[len(f.docs)-1] = crawler.Document{}
			f.docs = f.docs[:len(f.docs)-1]
		}
	}

	if !isDeleted {
		return fmt.Errorf("document not found")
	}
	return nil
}

func (f *FakeIndex) Update(ID uint64, url, title string) error {
	var isUpdated bool = false

	for i, doc := range f.docs {
		if doc.ID == ID {
			isUpdated = true
			f.docs[i].Title = title
			f.docs[i].URL = url
		}
	}

	if !isUpdated {
		return fmt.Errorf("document not found")
	}
	return nil
}

func (f *FakeIndex) Index() index.Index {
	return index.Index{"как": []uint64{1123123, 12432343, 1242544}}
}

func (f *FakeIndex) Find(request string) ([]crawler.Document, error) {
	var docs []crawler.Document

	for _, doc := range f.docs {
		if strings.Contains(doc.Title, request) {
			docs = append(docs, doc)
		}
	}

	if len(docs) == 0 {
		return nil, fmt.Errorf("document not found")
	}
	return docs, nil
}
