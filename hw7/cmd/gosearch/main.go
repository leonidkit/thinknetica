package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"gosearch/pkg/engine"
	"gosearch/pkg/index/invert"
)

type Gosearch struct {
	crawler crawler.Interface
	engine  *engine.Service
}

func ScanAsync(app *Gosearch, chdata chan<- []crawler.Document, url string, depth int, filename string) {
	data, err := app.crawler.Scan(url, depth)
	if err != nil {
		log.Printf("ошибка при получении данных с сайта %s: %v\n", url, err)
	}

	err = DumpFile(data, filename)
	if err != nil {
		log.Printf("ошибка при сохрании результатов сканирования: %v\n", err)
	}

	chdata <- data
}

func DumpFile(data []crawler.Document, filename string) error {
	buf := new(bytes.Buffer)
	e := gob.NewEncoder(buf)

	prepdata := make(map[string]string, len(data))
	for _, row := range data {
		prepdata[row.URL] = row.Title
	}

	err := e.Encode(prepdata)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, buf.Bytes(), 0664)
	if err != nil {
		return err
	}
	return nil
}

func LoadFile(filename string) ([]crawler.Document, error) {
	var data = make(map[string]string)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	bReader := bytes.NewReader(content)
	d := gob.NewDecoder(bReader)

	err = d.Decode(&data)
	if err != nil {
		return nil, err
	}

	var resdata []crawler.Document
	for url, title := range data {
		resdata = append(resdata, crawler.Document{
			URL:   url,
			Title: title,
		})
	}

	return resdata, nil
}

func main() {
	const url = "https://habr.com"
	const datafname = "data.gob"
	chdata := make(chan []crawler.Document)

	data, err := LoadFile(datafname)
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

	go ScanAsync(app, chdata, url, 1, datafname)

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
