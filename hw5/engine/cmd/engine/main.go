package main

import (
	"bufio"
	"engine/pkg/index"
	"engine/pkg/spider"
	"fmt"
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

func main() {
	const url = "https://habr.com"

	var cr = spider.New()
	data, err := cr.Scan(url, 1)
	if err != nil {
		log.Fatalf("ошибка при получении данных с сайта %s: %v\n", url, err)
	}

	var indexer = index.NewIndexTree(data)

	enter := "Enter word to find: "
	snr := bufio.NewScanner(os.Stdin)

	for fmt.Print(enter); snr.Scan(); fmt.Print(enter) {
		word := snr.Text()
		if strings.Replace(word, " ", "", -1) == "exit" {
			break
		}
		if word != "" {
			r, err := Found(indexer, word)
			if err != nil {
				log.Print(err.Error())
			}
			fmt.Print(r)
		}
	}

	if err := snr.Err(); err != nil {
		log.Fatal(err)
	}
}
