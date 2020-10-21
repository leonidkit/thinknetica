package main

import (
	"bufio"
	"bytes"
	"crawler/pkg/spider"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
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
	fsize := int64(0)
	_, err := os.Open(dataFilename)
	if err == nil {
		fStat, err := os.Stat(dataFilename)
		if err != nil {
			fsize = 0
		}
		fsize = fStat.Size()
	}

	if resetData || fsize == 0 {
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

func printFounded(data map[string]string, word string) {
	for k, v := range data {
		if strings.Contains(strings.ToLower(v), strings.ToLower(word)) {
			fmt.Printf("%s - %s\n", k, v)
		}
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

	data = getData(url, resetData, dataFile)

	if wordFind != "" {
		printFounded(data, wordFind)
	} else {
		enter := "Enter word to find: "
		snr := bufio.NewScanner(os.Stdin)

		for fmt.Print(enter); snr.Scan(); fmt.Print(enter) {
			word := snr.Text()
			if strings.Replace(word, " ", "", -1) == "exit" {
				break
			}
			if word != "" {
				printFounded(data, word)
			}
		}

		if err := snr.Err(); err != nil {
			if err != io.EOF {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}

}