package engine

import (
	"fmt"
	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
)

type Service struct {
	Tree index.Interface
	Data []crawler.Document
}

func New(index index.Interface, data []crawler.Document) *Service {
	return &Service{
		Tree: index,
		Data: data,
	}
}

func (s *Service) Search(query string) ([]crawler.Document, error) {
	if query == "" {
		return nil, fmt.Errorf("пустой запрос")
	}
	docs, err := s.Tree.Find(query)
	if err != nil {
		return nil, err
	}
	return docs, nil
}
