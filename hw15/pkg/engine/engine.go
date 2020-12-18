package engine

import (
	"fmt"
	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
)

type Service struct {
	Index index.Interface
	Data  []crawler.Document
}

func New(index index.Interface, data []crawler.Document) *Service {
	return &Service{
		Index: index,
		Data:  data,
	}
}

func (s *Service) Search(query string) ([]crawler.Document, error) {
	if query == "" {
		return nil, fmt.Errorf("пустой запрос")
	}
	docs, err := s.Index.Find(query)
	if err != nil {
		return nil, err
	}
	return docs, nil
}
