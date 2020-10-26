package main

import (
	"bufio"
	"bytes"
	"crawler/pkg/spider"
	"crawler/pkg/spiderblank"
	"encoding/gob"
	"flag"
	"fmt"
	"indexer/pkg/index"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Scanner interface {
	Scan(url string, depth int) (map[string]string, error)
}

type Crawler struct{}

func (c *Crawler) Scan(url string, depth int) (map[string]string, error) {
	return spider.Scan(url, depth)
}

type Crawlerblank struct{}

func (c *Crawlerblank) Scan(url string, depth int) (map[string]string, error) {
	return spiderblank.Scan(url, depth)
}

// Функция читает файл и возвращает декодированный из формата gob словарь.
func readFile(filename string) map[string]string {
	var data = make(map[string]string)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	bReader := bytes.NewReader(content)
	d := gob.NewDecoder(bReader)

	d.Decode(&data)

	return data
}

// Функция проверяет на возможность чтения файл dataFileName и если не указан флаг -r или -reset читает из этого файла.
// В противном случае, если файл невозможно прочитать или укзааны флаги -r и -reset сканирует сайт по url и записывает эти данные в dataFileName.
func getData(scanner Scanner, url string, resetData bool, dataFilename string) map[string]string {
	fsize := int64(0)
	_, err := os.Open(dataFilename)
	if err == nil {
		fStat, err := os.Stat(dataFilename)
		if err != nil {
			fsize = 0
		}
		fsize = fStat.Size()
	}

	if !resetData && fsize != 0 {
		return readFile(dataFilename)
	}

	data, err := scanner.Scan(url, 2)
	if err != nil {
		log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}
	defer func() {
		b := new(bytes.Buffer)
		e := gob.NewEncoder(b)
		err := e.Encode(data)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(dataFilename, b.Bytes(), 0664)
	}()

	return data
}

// Функция для происхождения по словарю data и вывода на печать совпадений с word.
func printFounded(i *index.Invert, word string) {
	recs := i.FindRecord(word)
	for _, rec := range recs {
		fmt.Printf("%s - %s\n", rec.URL, rec.Title)
	}
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

	var cr = &Crawler{}
	var sc Scanner = cr

	data = getData(sc, url, resetData, dataFile)

	var indexer = index.NewIndex(data)

	if wordFind != "" {
		printFounded(indexer, wordFind)
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
			printFounded(indexer, word)
		}
	}

	if err := snr.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
