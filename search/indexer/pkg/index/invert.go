package index

import (
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
			ID:    getHashNum(url),
			Title: title,
			URL:   url,
		}
		ldocs = append(ldocs, *newDoc)

		for _, word := range strings.Split(title, " ") {
			trimWord := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
				return unicode.IsPunct(r)
			}))
			lindx[trimWord] = append(lindx[trimWord], newDoc.ID)
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
		index := BinarySearch(address, i.docs)
		records = append(records, i.docs[index])
	}
	return records
}

// Ищет запись uint64 в слайсе []Document по значению, являющемся полем ID в структуре Document
func BinarySearch(value uint64, source []Document) int64 {
	left := int64(0)
	right := int64(len(source))

	for left <= right {
		mid := int64((left + right) / 2)

		if value == source[mid].ID {
			return mid
		}
		if value < source[mid].ID {
			right = mid
		} else {
			left = mid
		}
	}
	return -1
}

// Возвращает хэш от строки
func getHashNum(data string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(data))
	return h.Sum64()
}
