package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"gosearch/pkg/engine"
	"gosearch/pkg/index/invert"
	"gosearch/pkg/storage"
)

type Gosearch struct {
	crawler crawler.Interface
	engine  *engine.Service
}

func (g *Gosearch) ScanAsync(chdata chan<- []crawler.Document, flr *storage.Filer, url string, depth int, filename string) {
	data, err := g.crawler.Scan(url, depth)
	if err != nil {
		log.Fatalf("ошибка при получении данных с сайта %s: %v\n", url, err)
	}

	err = flr.DumpFile(data, filename)
	if err != nil {
		log.Fatalf("ошибка при сохрании результатов сканирования: %v\n", err)
	}

	chdata <- data
}

func main() {
	const url = "https://habr.com"
	const datafname = "data.gob"
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("не удалось получить абсолютный путь к файлу с данными")
	}

	chdata := make(chan []crawler.Document)
	flr := storage.New()

	log.Print(filepath.Join(path, datafname))
	data, err := flr.LoadFile(filepath.Join(path, datafname))
	if err != nil {
		log.Printf("ошибка при загрузки данных из файла %s: %v\n", datafname, err)
	}

	indx := invert.NewIndexTree(data)
	engn := engine.New(indx)
	spdr := spider.New()

	app := &Gosearch{
		crawler: spdr,
		engine:  engn,
	}

	//	go app.ScanAsync(chdata, flr, url, 2, datafname)

	enter := "Enter word to find: "
	snr := bufio.NewScanner(os.Stdin)

OUT:
	for fmt.Print(enter); snr.Scan(); fmt.Print(enter) {
		select {
		case data = <-chdata:
			app.engine.Index = invert.NewIndexTree(data)
		default:
			word := snr.Text()
			if strings.Replace(word, " ", "", -1) == "exit" {
				break OUT
			}
			if word != "" {
				recs, err := app.engine.Search(word)
				if err != nil {
					log.Printf("ошибка при поиске запроса %s: %v\n", word, err)
					continue OUT
				}
				for _, rec := range recs {
					fmt.Printf("%s - %s\n", rec.URL, rec.Title)
				}
			}
		}
	}

	if err := snr.Err(); err != nil {
		log.Fatal(err)
	}
}
