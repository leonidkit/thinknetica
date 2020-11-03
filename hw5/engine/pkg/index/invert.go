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
	Id    uint64
	Title string
	URL   string
}

type Index map[string][]uint64

type InvertList struct {
	index Index
	docs  []Document
}

type InvertTree struct {
	index Index
	docs  *btree.Tree
}

func (d *Document) ID() uint64 {
	return d.Id
}

// Возвращает новый объект структуры индекса на деревьях
func NewIndexTree(data map[string]string) *InvertTree {
	var tree = btree.NewTree()
	var lindx = make(Index, len(data)*3)

	for url, title := range data {
		newDoc := &Document{
			Id:    rand.Uint64(),
			Title: title,
			URL:   url,
		}
		tree.Add(newDoc)

		for _, word := range strings.Split(title, " ") {
			trimWord := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
				return unicode.IsPunct(r)
			}))
			if !checkOccurrence(lindx[trimWord], newDoc.Id) {
				lindx[trimWord] = append(lindx[trimWord], newDoc.Id)
			}
		}
	}

	return &InvertTree{
		lindx,
		tree,
	}
}

// Возвращает новый объект структуры индекса на списках
func NewIndexList(data map[string]string) *InvertList {
	var ldocs = make([]Document, 0, len(data))
	var lindx = make(Index, len(data)*3)

	// TODO: каждую запись в горутине обрабатывать и, соответственно, использовать sync.Map
	for url, title := range data {
		newDoc := &Document{
			Id:    rand.Uint64(),
			Title: title,
			URL:   url,
		}
		ldocs = append(ldocs, *newDoc)

		for _, word := range strings.Split(title, " ") {
			trimWord := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
				return unicode.IsPunct(r)
			}))
			if !checkOccurrence(lindx[trimWord], newDoc.Id) {
				lindx[trimWord] = append(lindx[trimWord], newDoc.Id)
			}
		}
	}

	sort.Slice(ldocs, func(i, j int) bool {
		return ldocs[i].Id < ldocs[j].Id
	})

	return &InvertList{
		lindx,
		ldocs,
	}
}

// Возвращает слайс []Document с записями в которых найдена strings
func (i *InvertList) FindRecord(record string) ([]Document, error) {
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
func (i *InvertTree) FindRecord(record string) ([]*Document, error) {
	records := []*Document{}
	docs := i.index[record]

	for _, address := range docs {
		res, err := i.docs.Search(&Document{
			Id:    address,
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

// Ищет запись uint64 в слайсе []Document по значению, являющемся полем Id в структуре Document
func binarySearch(value uint64, source []Document) (int64, error) {
	left := int64(0)
	right := int64(len(source))

	for left <= right {
		mId := int64((left + right) / 2)

		if value == source[mId].Id {
			return mId, nil
		}
		if value < source[mId].Id {
			right = mId - 1
		} else {
			left = mId + 1
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
