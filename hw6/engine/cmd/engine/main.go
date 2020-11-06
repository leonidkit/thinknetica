package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"engine/pkg/index"
	"engine/pkg/spider"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Scanner interface {
	Scan(url string, depth int) (map[string]string, error)
}

func Found(i *index.InvertedTree, word string) (string, error) {
	recs, err := i.FindRecord(word)
	if err != nil {
		return "", err
	}
	for _, rec := range recs {
		return fmt.Sprintf("%s - %s\n", rec.URL, rec.Title), nil
	}
	return "", nil
}

func DumpFile(data map[string]string, filename string) error {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)

	err := e.Encode(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b.Bytes(), 0664)
	if err != nil {
		return err
	}
	return nil
}

func LoadFile(filename string) (map[string]string, error) {
	var data = make(map[string]string)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return map[string]string{}, err
	}

	bReader := bytes.NewReader(content)
	d := gob.NewDecoder(bReader)

	d.Decode(&data)

	return data, nil
}

func main() {
	const url = "https://habr.com"
	const datafname = "data.gob"
	chdata := make(chan map[string]string)

	data, err := LoadFile(datafname)
	if err != nil {
		log.Printf("ошибка при загрузки данных из файла %s: %v\n", datafname, err)
	}

	go func(chdata chan<- map[string]string) {
		var cr = spider.New()
		data, err = cr.Scan(url, 2)
		if err != nil {
			log.Fatalf("ошибка при получении данных с сайта %s: %v\n", url, err)
		}

		err = DumpFile(data, datafname)
		if err != nil {
			log.Fatalf("ошибка при сохрании результатов сканирования: %v\n", err)
		}

		chdata <- data
	}(chdata)

	var indexer = index.NewIndexTree(data)

	enter := "Enter word to find: "
	snr := bufio.NewScanner(os.Stdin)

LOOP:
	for fmt.Print(enter); snr.Scan(); fmt.Print(enter) {
		select {
		case data = <-chdata:
			indexer = index.NewIndexTree(data)
		default:
			word := snr.Text()
			if strings.Replace(word, " ", "", -1) == "exit" {
				break LOOP
			}
			if word != "" {
				r, err := Found(indexer, word)
				if err != nil {
					log.Print(err.Error())
				}
				fmt.Print(r)
			}
		}
	}

	if err := snr.Err(); err != nil {
		log.Fatal(err)
	}
}
