package main

import (
	"fmt"
	"indexer/pkg/index"
)

var data = map[string]string{
	"http://habr.ru/main":        "Главная",
	"http://habr.ru/contact":     "Контакты",
	"http://habr.ru/comment":     "Комментарии ,Главная",
	"http://habr.ru/commentmain": "Комментарии ,Главная",
	"http://habr.ru/lol":         "Лол",
}

func main() {
	i := index.NewIndex(data)

	res := i.FindRecord("главная")

	fmt.Print(res)

}
