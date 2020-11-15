package storage

import (
	"bytes"
	"encoding/gob"
	"gosearch/pkg/crawler"
	"io/ioutil"
)

type Filer struct{}

func New() *Filer {
	return &Filer{}
}

// Серриализует переданные документы формата crawler'a в файл в формате gob
func (f *Filer) DumpFile(data []crawler.Document, filename string) error {
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

// Дессериализует документы формата crawler'a из файла
func (f *Filer) LoadFile(filename string) ([]crawler.Document, error) {
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
