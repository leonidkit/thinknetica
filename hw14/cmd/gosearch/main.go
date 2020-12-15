package main

import (
	"log"
	"sync"

	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"gosearch/pkg/engine"
	"gosearch/pkg/index/inverted"
	"gosearch/pkg/netsrv"
)

type Gosearch struct {
	crawler crawler.Interface
	engine  *engine.Service
}

func (g *Gosearch) ScanAsync(urls []string, depth int) {
	dataCh, errCh := g.crawler.BatchScan(urls, depth, 10)
	var data []crawler.Document

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for err := range errCh {
			log.Printf("ошибка при получении данных с сайта: %s\n", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		for doc := range dataCh {
			data = append(data, doc)
		}
	}()
	wg.Wait()

	indx := inverted.NewIndexTree(data)
	g.engine = engine.New(indx)
}

func main() {
	urls := []string{"https://habr.com", "https://go.dev", "https://golang.org/"}

	spdr := spider.New()
	app := &Gosearch{
		crawler: spdr,
	}
	app.ScanAsync(urls, 1)

	srv := netsrv.New("0.0.0.0", "8000", *app.engine)

	log.Fatal(srv.Serve())
}