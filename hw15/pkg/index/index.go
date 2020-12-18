package index

import "gosearch/pkg/crawler"

type Index map[string][]uint64

type Interface interface {
	Find(string) ([]crawler.Document, error)
	Recieve() Index
}
