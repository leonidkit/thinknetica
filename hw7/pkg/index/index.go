package index

import "gosearch/pkg/crawler"

type Interface interface {
	Find(string) ([]crawler.Document, error)
}
