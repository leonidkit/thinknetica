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
	"gosearch/pkg/index/inverted"
	"gosearch/pkg/storage"
)

type Gosearch struct {
	crawler crawler.Interface
	engine  *engine.Service
	filer   *storage.Filer
}

func (g *Gosearch) ScanAsync(url string, depth int, filename string) {
	data, err := g.crawler.Scan(url, depth)
	if err != nil {
		log.Fatalf("ошибка при получении данных с сайта %s: %v\n", url, err)
	}

	err = g.filer.DumpFile(data, filename)
	if err != nil {
		log.Fatalf("ошибка при сохрании результатов сканирования: %v\n", err)
	}

	g.engine.Index = inverted.NewIndexTree(data)
}

func main() {
	const url = "https://habr.com"
	const datafname = "data.gob"
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("не удалось получить абсолютный путь к файлу с данными")
	}

	spdr := spider.New()
	flr := storage.New()

	app := &Gosearch{
		crawler: spdr,
		filer:   flr,
	}

	data, err := app.filer.LoadFile(filepath.Join(path, datafname))
	if err != nil {
		log.Printf("ошибка при загрузки данных из файла %s: %v\n", datafname, err)
	}

	indx := inverted.NewIndexTree(data)
	app.engine = engine.New(indx)

	go app.ScanAsync(url, 2, datafname)

	enter := "Enter word to find: "
	snr := bufio.NewScanner(os.Stdin)

	for fmt.Print(enter); snr.Scan(); fmt.Print(enter) {
		word := snr.Text()
		if strings.Replace(word, " ", "", -1) == "exit" {
			break
		}
		if word != "" {
			recs, err := app.engine.Search(word)
			if err != nil {
				log.Printf("ошибка при поиске запроса %s: %v\n", word, err)
				continue
			}
			for _, rec := range recs {
				fmt.Printf("%s - %s\n", rec.URL, rec.Title)
			}
		}
	}

	if err := snr.Err(); err != nil {
		log.Fatal(err)
	}
}
