package fakebot

type Scan struct{}

func New() *Scan {
	return &Scan{}
}

func (l *Scan) Scan(url string, depth int) (data map[string]string, err error) {
	data = map[string]string{
		"https://habr.ru":         "Главная",
		"https://habr.ru/contact": "Контакты",
	}

	return data, nil
}
