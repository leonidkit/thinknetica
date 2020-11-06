// Содержит стурктуру, представляющую собой инвертированный индекс и методы этой структуры
// для создания индекса и поиска по этому индексу
package index

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"unicode"

	"engine/pkg/index/btree"
)

type Document struct {
	ID    uint64
	Title string
	URL   string
}

type Index map[string][]uint64

type InvertedList struct {
	index Index
	docs  []Document
}

type InvertedTree struct {
	index Index
	docs  *btree.Tree
}

func (d *Document) Ident() uint64 {
	return d.ID
}

func NewIndexTree(data map[string]string) *InvertedTree {
	var tree = btree.New()
	var lindx = make(Index, len(data)*3)

	for url, title := range data {
		newDoc := &Document{
			ID:    rand.Uint64(),
			Title: title,
			URL:   url,
		}
		tree.Add(newDoc)

		for _, word := range strings.Split(title, " ") {
			trimWord := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
				return unicode.IsPunct(r)
			}))
			if !checkOccurrence(lindx[trimWord], newDoc.ID) {
				lindx[trimWord] = append(lindx[trimWord], newDoc.ID)
			}
		}
	}

	return &InvertedTree{
		lindx,
		tree,
	}
}

func NewIndexList(data map[string]string) *InvertedList {
	var ldocs = make([]Document, 0, len(data))
	var lindx = make(Index, len(data)*3)

	// TODO: каждую запись в горутине обрабатывать и, соответственно, использовать sync.Map
	for url, title := range data {
		newDoc := &Document{
			ID:    rand.Uint64(),
			Title: title,
			URL:   url,
		}
		ldocs = append(ldocs, *newDoc)

		for _, word := range strings.Split(title, " ") {
			trimWord := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
				return unicode.IsPunct(r)
			}))
			if !checkOccurrence(lindx[trimWord], newDoc.ID) {
				lindx[trimWord] = append(lindx[trimWord], newDoc.ID)
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

// Возвращает слайс []Document с записями в которых найдена strings
func (i *InvertedList) FindRecord(record string) ([]Document, error) {
	records := []Document{}
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

// Возвращает слайс []Document с записями в которых найдена strings
func (i *InvertedTree) FindRecord(record string) ([]*Document, error) {
	records := []*Document{}
	docs := i.index[record]

	for _, address := range docs {
		res, err := i.docs.Search(&Document{
			ID:    address,
			Title: "",
			URL:   "",
		})
		if err != nil {
			return nil, err
		}
		records = append(records, res.(*Document))
	}
	return records, nil
}

// Ищет запись uint64 в слайсе []Document по значению, являющемся полем ID в структуре Document
func binarySearch(value uint64, source []Document) (int64, error) {
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
