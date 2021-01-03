package index

import (
	"gosearch/pkg/crawler"
)

type Index map[string][]uint64

type Interface interface {
	Find(string) ([]crawler.Document, error)
	Index() Index
	Add(string, string) error
	Delete(uint64) error
	Update(uint64, string, string) error
}
