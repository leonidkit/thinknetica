package main

import (
	"bytes"
	"crawler/pkg/spider"
	"encoding/gob"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

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

func getData(url string, resetData bool, dataFilename string) (data map[string]string) {
	if resetData {
		data, err := spider.Scan(url, 2)
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
	} else {
		data = readFile(dataFilename)
	}
	return
}

func main() {
	const url = "https://habr.com"
	var dataFile string = "data.gob"
	var data = make(map[string]string)

	var resetData bool
	var wordFind string

	flag.BoolVar(&resetData, "reset", false, "saved фмув data will be deleted")
	flag.BoolVar(&resetData, "r", false, "saved data will be deleted")

	flag.StringVar(&wordFind, "w", "", "the word to be searched for")
	flag.StringVar(&wordFind, "word", "", "the word to be searched for")
	flag.Parse()

	if wordFind == "" {
		fmt.Println("Не задано слова для поиска")
		return
	}

	data = getData(url, resetData, dataFile)

	for k, v := range data {
		if strings.Contains(strings.ToLower(v), strings.ToLower(wordFind)) {
			fmt.Printf("%s - %s\n", k, v)
		}
	}
}
