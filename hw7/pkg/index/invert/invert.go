// Содержит стурктуру, представляющую собой инвертированный индекс и методы этой структуры
// для создания индекса и поиска по этому индексу
package invert

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"unicode"

	"gosearch/pkg/crawler"
	"gosearch/pkg/index/invert/btree"
)

type Index map[string][]uint64

type InvertedList struct {
	index Index
	docs  []crawler.Document
}

type InvertedTree struct {
	index Index
	docs  *btree.Tree
}

func NewIndexTree(data []crawler.Document) *InvertedTree {
	var tree = btree.New()
	var lindx = make(Index, len(data)*3)

	for _, row := range data {
		row.ID = rand.Uint64()
		tree.Add(&row)

		for _, word := range strings.Split(row.Title, " ") {
			trimWord := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
				return unicode.IsPunct(r)
			}))
			if !checkOccurrence(lindx[trimWord], row.ID) {
				lindx[trimWord] = append(lindx[trimWord], row.ID)
			}
		}
	}

	return &InvertedTree{
		lindx,
		tree,
	}
}

func NewIndexList(data []crawler.Document) *InvertedList {
	var ldocs = make([]crawler.Document, 0, len(data))
	var lindx = make(Index, len(data)*3)

	// TODO: каждую запись в горутине обрабатывать и, соответственно, использовать sync.Map
	for _, row := range data {
		row.ID = rand.Uint64()
		ldocs = append(ldocs, row)

		for _, word := range strings.Split(row.Title, " ") {
			trimWord := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
				return unicode.IsPunct(r)
			}))
			if !checkOccurrence(lindx[trimWord], row.ID) {
				lindx[trimWord] = append(lindx[trimWord], row.ID)
			}
		}
	}

	sort.Slice(ldocs, func(i, j int) bool {
		return ldocs[i].ID < ldocs[j].ID
	})

	return &InvertedList{
		lindx,
		ldocs,
	}
}

// Возвращает слайс []crawler.Document с записями в которых найдена strings
func (i *InvertedList) Find(record string) ([]crawler.Document, error) {
	records := []crawler.Document{}
	docs := i.index[record]

	for _, address := range docs {
		index, err := binarySearch(address, i.docs)
		if err != nil {
			return nil, err
		}
		records = append(records, i.docs[index])
	}
	return records, nil
}

// Возвращает слайс []crawler.Document с записями в которых найдена strings
func (i *InvertedTree) Find(record string) ([]*crawler.Document, error) {
	records := []*crawler.Document{}
	docs := i.index[record]

	for _, address := range docs {
		res, err := i.docs.Search(&crawler.Document{
			ID:    address,
			Title: "",
			URL:   "",
		})
		if err != nil {
			return nil, err
		}
		records = append(records, res.(*crawler.Document))
	}
	return records, nil
}

// Ищет запись uint64 в слайсе []crawler.Document по значению, являющемся полем ID в структуре crawler.Document
func binarySearch(value uint64, source []crawler.Document) (int64, error) {
	left := int64(0)
	right := int64(len(source))

	for left <= right {
		mID := int64((left + right) / 2)

		if value == source[mID].ID {
			return mID, nil
		}
		if value < source[mID].ID {
			right = mID - 1
		} else {
			left = mID + 1
		}
	}
	return int64(0), fmt.Errorf("nothing found")
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
