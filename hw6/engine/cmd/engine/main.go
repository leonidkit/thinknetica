package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"engine/pkg/index"
	"engine/pkg/spider"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Scanner interface {
	Scan(url string, depth int) (map[string]string, error)
}

func DumpFile(buf io.Writer, data map[string]string, filename string) error {
	e := gob.NewEncoder(buf)

	err := e.Encode(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, buf.Bytes(), 0664)
	if err != nil {
		return err
	}
	return nil
}

func LoadFile(filename string) (map[string]string, error) {
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

	return data, nil
}

func ScanAsync(chdata chan<- map[string]string, url string, filename string) {
	var cr = spider.New()
	data, err := cr.Scan(url, 2)
	if err != nil {
		log.Printf("ошибка при получении данных с сайта %s: %v\n", url, err)
	}

	buf := new(bytes.Buffer)
	err = DumpFile(buf, data, filename)
	if err != nil {
		log.Printf("ошибка при сохрании результатов сканирования: %v\n", err)
	}

	chdata <- data
}

func main() {
	const url = "https://habr.com"
	const datafname = "data.gob"
	chdata := make(chan map[string]string)

	data, err := LoadFile(datafname)
	if err != nil {
		log.Fatalf("ошибка при загрузки данных из файла %s: %v\n", datafname, err)
	}

	go ScanAsync(chdata, url, datafname)

	var indexer = index.NewIndexTree(data)

	enter := "Enter word to find: "
	snr := bufio.NewScanner(os.Stdin)

OUT:
	for fmt.Print(enter); snr.Scan(); fmt.Print(enter) {
		select {
		case data = <-chdata:
			indexer = index.NewIndexTree(data)
		default:
			word := snr.Text()
			if strings.Replace(word, " ", "", -1) == "exit" {
				break OUT
			}
			if word != "" {
				recs, err := indexer.FindRecord(word)
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
