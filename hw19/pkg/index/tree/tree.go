// Содержит стурктуру, представляющую собой инвертированный индекс и методы этой структуры
// для создания индекса и поиска по этому индексу
package tree

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"

	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
	"gosearch/pkg/index/tree/btree"
)

type Tree struct {
	index index.Index
	docs  *btree.Tree
}

func NewTree(data []crawler.Document) *Tree {
	var tree = btree.New()
	var lindx = make(index.Index, len(data)*3)

	for _, row := range data {
		row.ID = rand.Uint64()
		tree.Add(row)

		for _, word := range strings.Split(row.Title, " ") {
			trimWord := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
				return unicode.IsPunct(r)
			}))

			if !checkOccurrence(lindx[trimWord], row.ID) && trimWord != "" {
				lindx[trimWord] = append(lindx[trimWord], row.ID)
			}
		}
	}

	return &Tree{
		lindx,
		tree,
	}
}

func (i *Tree) Add(url string, title string) error {
	var row crawler.Document
	row.ID = rand.Uint64()
	row.Title = title
	row.URL = url

	err := i.docs.Add(row)
	if err != nil {
		return nil
	}

	for _, word := range strings.Split(row.Title, " ") {
		trimWord := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
			return unicode.IsPunct(r)
		}))

		if !checkOccurrence(i.index[trimWord], row.ID) && trimWord != "" {
			i.index[trimWord] = append(i.index[trimWord], row.ID)
		}
	}

	return nil
}

func (i *Tree) Delete(ID uint64) error {
	row := crawler.Document{
		ID: ID,
	}

	tmp, err := i.docs.Delete(row)
	if err != nil {
		return err
	}
	i.docs = tmp

	return nil
}

func (i *Tree) Update(ID uint64, url, title string) error {
	row := crawler.Document{
		ID:    ID,
		URL:   url,
		Title: title,
	}

	err := i.docs.Update(row)
	if err != nil {
		return err
	}

	return nil
}

func (i *Tree) Index() index.Index {
	return i.index
}

// Возвращает слайс []crawler.Document с записями в которых найдена strings
func (i *Tree) Find(record string) ([]crawler.Document, error) {
	records := []crawler.Document{}
	docs, exist := i.index[record]
	if !exist {
		return records, fmt.Errorf("index not found")
	}

	for _, address := range docs {
		res, err := i.docs.Search(&crawler.Document{
			ID:    address,
			Title: "",
			URL:   "",
		})
		if err != nil {
			return nil, err
		}
		records = append(records, res.(crawler.Document))
	}
	return records, nil
}

// Ищет запись uint64 в слайсе []crawler.Document по значению, являющемся полем ID в структуре crawler.Document
func binarySearch(value uint64, source []crawler.Document) (uint64, error) {
	left := uint64(0)
	right := uint64(len(source) - 1)

	for left <= right {
		mID := uint64((left + right) / 2)

		if value == source[mID].ID {
			return mID, nil
		}
		if value < source[mID].ID {
			right = mID - 1
		} else {
			left = mID + 1
		}
	}
	return uint64(0), fmt.Errorf("nothing found")
}

// Возвращает true если value имеется в слайсе data, иначе false
func checkOccurrence(data []uint64, value uint64) bool {
	for _, val := range data {
		if val == value {
			return true
		}
	}
	return false
}
