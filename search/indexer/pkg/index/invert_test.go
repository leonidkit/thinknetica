package index

import "testing"

func BenchmarkNewIndex(b *testing.B) {
	var index = &Index{}
	for i := 0; i < b.N; i++ {
		var data = map[string]string{
			"http://habr.ru/main":    "Главная",
			"http://habr.ru/contact": "Контакты",
			"http://habr.ru/comment": "Комментарии",
			"http://habr.ru/lol":     "Лол",
		}
		index.NewIndex(data)
	}
}
