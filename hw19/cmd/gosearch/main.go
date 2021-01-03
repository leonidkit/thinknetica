package main

import (
	"log"
	"net/http"
	"sync"

	"gosearch/pkg/api"
	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"gosearch/pkg/engine"
	"gosearch/pkg/index/tree"
	"gosearch/pkg/rpcsrv"

	"github.com/gorilla/mux"
)

type gosearch struct {
	crawler crawler.Interface
	engine  *engine.Service
	api     *api.Service
	rpcsrv  *rpcsrv.RPCsrv
}

func (g *gosearch) scanAsync(urls []string, depth int) {
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

	tree := tree.NewTree(data)
	g.engine = engine.New(tree, data)
	g.api = api.New(mux.NewRouter(), g.engine)
}

func main() {
	urls := []string{"https://habr.com", "https://go.dev", "https://golang.org/"}
	port := "8000"
	rpcport := "8001"
	host := "0.0.0.0"

	tree := tree.NewTree([]crawler.Document{})

	spdr := spider.New()
	engn := engine.New(tree, []crawler.Document{})

	app := &gosearch{
		crawler: spdr,
		engine:  engn,
		api:     api.New(mux.NewRouter(), engn),
	}

	app.scanAsync(urls, 1)
	go rpcsrv.Serve(host, rpcport, app.engine)

	http.ListenAndServe(host+":"+port, app.api.Router)
}
