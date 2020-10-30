package index

import (
	"fmt"
	"hash/fnv"
	"sort"
	"strings"
	"unicode"
)

type Document struct {
	ID    uint64
	Title string
	URL   string
}

type Index map[string][]uint64
type Docs []Document

type Invert struct {
	index Index
	docs  Docs
}

func NewIndex(data map[string]string) *Invert {
	var ldocs = make(Docs, 0, len(data))
	var lindx = make(Index, len(data)*3)

	// TODO: каждую запись в горутине обрабатывать и, соответственно, использовать sync.Map
	for url, title := range data {
		newDoc := &Document{
			ID:    hashNum(url),
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

	return &Invert{
		lindx,
		ldocs,
	}
}

// Возвращает слайс []Document с записями в которых найдена strings
func (i *Invert) FindRecord(record string) []Document {
	records := []Document{}
	docs := i.index[record]

	for _, address := range docs {
		index, err := BinarySearch(address, i.docs)
		if err != nil {
			continue
		}
		records = append(records, i.docs[index])
	}
	return records
}

// Ищет запись uint64 в слайсе []Document по значению, являющемся полем ID в структуре Document
func BinarySearch(value uint64, source []Document) (uint64, error) {
	left := uint64(0)
	right := uint64(len(source) - 1)

	for left <= right {
		mid := uint64((left + right) / 2)

		if value == source[mid].ID {
			return mid, nil
		}
		if value < source[mid].ID {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return uint64(0), fmt.Errorf("nothing found")
}

// Возвращает хэш от строки
func hashNum(data string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(data))
	return h.Sum64()
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
