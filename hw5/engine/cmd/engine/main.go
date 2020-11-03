package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"engine/pkg/index"
	"engine/pkg/spider"
	"flag"
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

// Функция читает файл и возвращает декодированный из формата gob словарь.
func readFile(filename string) (map[string]string, error) {
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

// Функция проверяет на возможность чтения файл dataFileName и если не указан флаг -r или -reset читает из этого файла.
// В противном случае, если файл невозможно прочитать или укзааны флаги -r и -reset сканирует сайт по url и записывает эти данные в dataFileName.
func getData(scanner Scanner, url string, resetData bool, dataFilename string) (map[string]string, error) {
	// чтение данных из файла
	if !resetData {
		// проверка на возможность чтения данных из файла
		_, err := os.Open(dataFilename)
		if err != nil {
			return map[string]string{}, err
		}

		data, err := readFile(dataFilename)
		if err != nil {
			return map[string]string{}, err
		}

		return data, nil
	}

	// чтение данных из сети
	data, err := scanner.Scan(url, 2)
	if err != nil {
		return map[string]string{}, err
	}

	// обновление данных в файле
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)

	err = e.Encode(data)
	if err != nil {
		return map[string]string{}, err
	}

	err = ioutil.WriteFile(dataFilename, b.Bytes(), 0664)
	if err != nil {
		return map[string]string{}, err
	}

	return data, nil
}

// Функция для происхождения по словарю data и вывода на печать совпадений с word.
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

	var dataFile string = "data.gob"
	var data = make(map[string]string)

	var resetData bool
	var wordFind string

	flag.BoolVar(&resetData, "reset", false, "saved data will be deleted")
	flag.BoolVar(&resetData, "r", false, "saved data will be deleted")

	flag.StringVar(&wordFind, "w", "", "the word to be searched for")
	flag.StringVar(&wordFind, "word", "", "the word to be searched for")
	flag.Parse()

	var cr = spider.NewSpider()
	data, err := getData(cr, url, resetData, dataFile)
	if err != nil {
		log.Fatalf("ошибка при получении данных с сайта %s: %v\n", url, err)
	}

	var indexer = index.NewIndexTree(data)

	if wordFind != "" {
		err = printFound(indexer, wordFind)
		if err != nil {
			log.Print(err.Error())
		}
		return
	}

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
