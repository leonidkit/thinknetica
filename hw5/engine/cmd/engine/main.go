package main

import (
	"bufio"
	"engine/pkg/index"
	"engine/pkg/spider"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Scanner interface {
	Scan(url string, depth int) (map[string]string, error)
}

func printFound(i *index.InvertTree, word string) error {
	recs, err := i.FindRecord(word)
	if err != nil {
		return err
	}
	for _, rec := range recs {
		fmt.Printf("%s - %s\n", rec.URL, rec.Title)
	}
	return nil
}

func main() {
	const url = "https://habr.com"

	var cr = spider.NewSpider()
	data, err := cr.Scan(url, 2)
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
			err = printFound(indexer, word)
			if err != nil {
				log.Print(err.Error())
			}
		}
	}

	if err := snr.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
